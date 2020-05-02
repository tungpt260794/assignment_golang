package services

import (
	"assignment/models"
	"time"

	"github.com/jinzhu/gorm"
)

// CreateGallery ...
func CreateGallery(db *gorm.DB, _gallery *models.Gallery) (gallery *models.Gallery, err error) {
	gallery = _gallery
	gallery.CreatedAt = time.Now()
	gallery.UpdatedAt = time.Now()

	err = db.Create(gallery).Error
	return
}

// UpdateGallery ...
func UpdateGallery(db *gorm.DB, id int, _gallery *models.Gallery) (err error) {

	gallery := &models.Gallery{}
	err = db.Where("id = ?", id).First(gallery).Error
	if err != nil {
		return
	}

	if _gallery.Name != "" {
		gallery.Name = _gallery.Name
	}
	if _gallery.Brief != "" {
		gallery.Brief = _gallery.Brief
	}
	if _gallery.Visibility != "" {
		gallery.Visibility = _gallery.Visibility
	}

	gallery.UpdatedAt = time.Now()

	err = db.Save(gallery).Error
	return
}

// GetGallery ...
func GetGallery(db *gorm.DB, id int) (gallery *models.Gallery, err error) {
	gallery = &models.Gallery{}
	err = db.First(gallery, id).Error
	return
}

// GetPublicGalleries ...
func GetPublicGalleries(db *gorm.DB) (galleries *[]models.Gallery, err error) {
	galleries = &[]models.Gallery{}
	err = db.Where("visibility = ?", "public").Find(galleries).Error
	if err != nil {
		return
	}

	for index, gallery := range *galleries {
		photos := []models.Photo{}
		photos, err = GetPhotosByGallery(db, gallery)
		if err != nil {
			return
		}
		(*galleries)[index].Photos = photos
	}

	return
}

// GetPublicGallery ...
func GetPublicGallery(db *gorm.DB, id int) (gallery *models.Gallery, err error) {
	gallery = &models.Gallery{}
	err = db.Where("visibility = ? AND id = ?", "public", id).First(gallery).Error
	if err != nil {
		return
	}

	photos := []models.Photo{}
	photos, err = GetPhotosByGallery(db, *gallery)
	if err != nil {
		return
	}
	gallery.Photos = photos

	for index, photo := range photos {
		reactions := []models.Reaction{}
		reactions, err = GetReactionsByPhoto(db, photo)
		if err != nil {
			return
		}
		photos[index].Reactions = reactions
	}

	return
}
