package routes

import (
	"assignment/models"
	"assignment/services"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateGallery ...
func CreateGallery(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	gallery := &models.Gallery{}
	err := ctx.BindJSON(gallery)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	gallery, err = services.CreateGallery(db, &models.Gallery{
		AccountID: accountID.(int),
		Name:      gallery.Name,
		Brief:     gallery.Brief,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, gallery)
	return
}

// UpdateGallery ...
func UpdateGallery(db *gorm.DB, ctx *gin.Context) {
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

	gallery, err := services.GetGallery(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if accountID.(int) != gallery.AccountID {
		ctx.AbortWithError(400, errors.New("YOU_NO_OWN_OF_GALLERY"))
		return
	}

	gallery = &models.Gallery{}
	err = ctx.BindJSON(gallery)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = services.UpdateGallery(db, int(id), &models.Gallery{
		Name:  gallery.Name,
		Brief: gallery.Brief,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// PublicGallery ...
func PublicGallery(db *gorm.DB, ctx *gin.Context) {
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

	gallery, err := services.GetGallery(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	if accountID.(int) != gallery.AccountID {
		ctx.AbortWithError(400, errors.New("YOU_NO_OWN_OF_GALLERY"))
		return
	}

	err = services.UpdateGallery(db, int(id), &models.Gallery{
		Visibility: "public",
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// GetPublicGalleries ...
func GetPublicGalleries(db *gorm.DB, ctx *gin.Context) {
	galleries := &[]models.Gallery{}

	galleries, err := services.GetPublicGalleries(db)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, galleries)
	return
}

// GetPublicGallery ...
func GetPublicGallery(db *gorm.DB, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	gallery, err := services.GetPublicGallery(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, gallery)
	return
}
