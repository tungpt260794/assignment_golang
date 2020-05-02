package services

import (
	"assignment/models"

	"github.com/jinzhu/gorm"
	// ...
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect ...
func Connect(connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Account{}, &models.Gallery{}, &models.Photo{}, &models.Reaction{}).Error
	if err != nil {
		return
	}

	db.Model(&models.Gallery{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")
	db.Model(&models.Photo{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")
	db.Model(&models.Photo{}).AddForeignKey("gallery_id", "galleries(id)", "CASCADE", "CASCADE")
	db.Model(&models.Reaction{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")
	db.Model(&models.Reaction{}).AddForeignKey("photo_id", "photos(id)", "CASCADE", "CASCADE")

	db.Model(&models.Account{}).AddUniqueIndex("unique_account_email", "email")

	return
}
