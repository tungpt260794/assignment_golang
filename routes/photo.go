package routes

import (
	"assignment/helpers"
	"assignment/models"
	"assignment/services"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UploadPhoto ...
func UploadPhoto(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	_galleryID := ctx.PostForm("galleryId")
	galleryID, err := strconv.ParseInt(_galleryID, 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	gallery, err := services.GetGallery(db, int(galleryID))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if accountID.(int) != gallery.AccountID {
		ctx.AbortWithError(400, errors.New("YOU_NO_OWN_OF_GALLERY"))
		return
	}

	path, size, err := helpers.UploadFile(ctx)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	width, _, err := helpers.GetDimension(path)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if width < 1024 {
		ctx.AbortWithError(400, errors.New("UPLOADED_PHOTO_MUST_HAS_WIDTH_GREATER_THAN_1024"))
		return
	}

	photo, err := services.CreatePhoto(db, &models.Photo{
		GalleryID:   gallery.ID,
		AccountID:   gallery.AccountID,
		Name:        name,
		Description: description,
		Size:        size,
		Path:        path,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, photo)
	return
}

// UpdatePhoto ...
func UpdatePhoto(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	photo, err := services.GetPhoto(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if accountID.(int) != photo.AccountID {
		ctx.AbortWithError(400, errors.New("YOU_NO_OWN_OF_PHOTO"))
		return
	}

	photo = &models.Photo{}
	err = ctx.BindJSON(photo)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = services.UpdatePhoto(db, int(id), &models.Photo{
		Name:        photo.Name,
		Description: photo.Description,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// Like ...
func Like(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	reaction, err := services.CreateReaction(db, &models.Reaction{
		AccountID: accountID.(int),
		PhotoID:   int(id),
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, reaction)
	return
}

// UnLike ...
func UnLike(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = services.DeleteReaction(db, accountID.(int), int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// GetPhoto ...
func GetPhoto(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	photo, err := services.GetPhoto(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if accountID.(int) != photo.AccountID {
		ctx.AbortWithError(400, errors.New("YOU_NO_OWN_OF_PHOTO"))
		return
	}

	reactions, err := services.GetReactionsByPhotoID(db, photo.ID)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	photo.Reactions = reactions
	ctx.JSON(200, photo)
	return
}

// DownloadPhoto ...
func DownloadPhoto(db *gorm.DB, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	width, err := strconv.ParseInt(ctx.Param("width"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	photo, err := services.GetPhoto(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	path := ""
	if width == 1920 {
		path = photo.W1920Path
	}
	if width == 1600 {
		path = photo.W1600Path
	}
	if width == 1280 {
		path = photo.W1280Path
	}
	if width == 1024 {
		path = photo.W1024Path
	}
	if width == 800 {
		path = photo.W800Path
	}
	if width == 256 {
		path = photo.W256Path
	}

	header := ctx.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + time.Now().Format("20060102150405") + ".jpg"}

	file, err := os.Open(path)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	defer file.Close()

	io.Copy(ctx.Writer, file)
	ctx.Status(200)
	return
}
