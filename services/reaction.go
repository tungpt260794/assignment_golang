package services

import (
	"assignment/models"
	"time"

	"github.com/jinzhu/gorm"
)

// CreateReaction ...
func CreateReaction(db *gorm.DB, _reaction *models.Reaction) (reaction *models.Reaction, err error) {
	reaction = _reaction
	reaction.CreatedAt = time.Now()
	reaction.UpdatedAt = time.Now()

	err = db.Create(reaction).Error
	return
}

// DeleteReaction ...
func DeleteReaction(db *gorm.DB, accountID int, photoID int) (err error) {
	err = db.Unscoped().Delete(&models.Reaction{AccountID: accountID, PhotoID: photoID}).Error
	return
}

// GetReactionsByPhotoID ...
func GetReactionsByPhotoID(db *gorm.DB, photoID int) (reactions []models.Reaction, err error) {
	reactions = []models.Reaction{}
	err = db.Where("photo_id = ?", photoID).Find(&reactions).Error
	return
}

// GetReactionsByPhoto ...
func GetReactionsByPhoto(db *gorm.DB, photo models.Photo) (reactions []models.Reaction, err error) {
	reactions = []models.Reaction{}
	err = db.Model(&photo).Related(&reactions).Error
	if err != nil {
		return
	}

	return
}
