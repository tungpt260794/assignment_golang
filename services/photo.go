package services

import (
	"assignment/helpers"
	"assignment/models"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// CreatePhoto ...
func CreatePhoto(db *gorm.DB, _photo *models.Photo) (photo *models.Photo, err error) {
	sizes, err := helpers.ResizeAll(_photo.Path)
	fmt.Println(sizes[1280])

	photo = _photo
	if sizes[1920] != "" {
		photo.W1920Path = sizes[1920]
	}
	if sizes[1600] != "" {
		photo.W1600Path = sizes[1600]
	}
	if sizes[1280] != "" {
		photo.W1280Path = sizes[1280]
	}
	if sizes[1024] != "" {
		photo.W1024Path = sizes[1024]
	}
	if sizes[800] != "" {
		photo.W800Path = sizes[800]
	}
	if sizes[256] != "" {
		photo.W256Path = sizes[256]
	}
	photo.CreatedAt = time.Now()
	photo.UpdatedAt = time.Now()

	err = db.Create(photo).Error
	return
}

// UpdatePhoto ...
func UpdatePhoto(db *gorm.DB, id int, _photo *models.Photo) (err error) {
	photo := &models.Photo{}
	err = db.Where("id = ?", id).First(photo).Error
	if err != nil {
		return
	}

	if _photo.Name != "" {
		photo.Name = _photo.Name
	}
	if _photo.Description != "" {
		photo.Description = _photo.Description
	}

	photo.UpdatedAt = time.Now()

	err = db.Save(photo).Error
	return
}

// GetPhoto ...
func GetPhoto(db *gorm.DB, id int) (photo *models.Photo, err error) {
	photo = &models.Photo{}
	err = db.First(photo, id).Error
	return
}

// GetPhotosByGallery ...
func GetPhotosByGallery(db *gorm.DB, gallery models.Gallery) (photos []models.Photo, err error) {
	photos = []models.Photo{}
	err = db.Model(&gallery).Related(&photos).Error
	if err != nil {
		return
	}

	return
}
