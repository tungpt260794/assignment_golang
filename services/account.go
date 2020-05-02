package services

import (
	"assignment/helpers"
	"assignment/models"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Login ...
func Login(db *gorm.DB, email string, password string) (tokenStr string, err error) {
	account := &models.Account{}
	err = db.Where("email = ?", email).First(account).Error
	if err != nil {
		return
	}

	if !CheckPasswordHash(password, account.Password) {
		err = errors.New("PASSWORD_INVALID")
		return
	}

	tokenStr, err = CreateToken(account.ID)
	return
}

// CreateAccount ...
func CreateAccount(db *gorm.DB, _account *models.Account) (account *models.Account, err error) {
	passwordHashed, err := HashPassword(_account.Password)
	if err != nil {
		return
	}

	account = _account
	account.Password = passwordHashed
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err = db.Create(account).Error
	return
}

// UpdateAccount ...
func UpdateAccount(db *gorm.DB, id int, _account *models.Account) (err error) {

	account := &models.Account{}
	err = db.Where("id = ?", id).First(account).Error
	if err != nil {
		return
	}

	if _account.Email != "" {
		account.Email = _account.Email
	}
	if _account.Password != "" {
		account.Password = _account.Password
	}
	if _account.Name != "" {
		account.Name = _account.Name
	}
	if _account.Phone != "" {
		account.Phone = _account.Phone
	}
	if _account.Address != "" {
		account.Address = _account.Address
	}
	if _account.Avatar != "" {
		src := ""
		src, err = helpers.Resize(_account.Avatar, 64, 64)
		if err != nil {
			return
		}

		account.Avatar = src
	}

	account.UpdatedAt = time.Now()

	err = db.Save(account).Error
	return
}

// GetAccount ...
func GetAccount(db *gorm.DB, id int) (account *models.Account, err error) {
	account = &models.Account{}
	err = db.First(account, id).Error
	return
}
