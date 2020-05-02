package routes

import (
	"assignment/helpers"
	"assignment/models"
	"assignment/services"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Login ...
func Login(db *gorm.DB, ctx *gin.Context) {
	account := &models.Account{}
	err := ctx.BindJSON(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	tokenStr, err := services.Login(db, account.Email, account.Password)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.String(200, tokenStr)
	return
}

// Register ...
func Register(db *gorm.DB, ctx *gin.Context) {
	account := &models.Account{}
	err := ctx.BindJSON(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	account, err = services.CreateAccount(db, &models.Account{
		Email:    account.Email,
		Password: account.Password,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, models.Account{
		ID:        account.ID,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	})
	return
}

// ChangePasswordStruct ...
type ChangePasswordStruct struct {
	PasswordOld     string `json:"password_old"`
	PasswordNew     string `json:"password_new"`
	PasswordConfirm string `json:"password_confirm"`
}

// ChangePassword ...
func ChangePassword(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	account, err := services.GetAccount(db, accountID.(int))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	changePasswordStruct := &ChangePasswordStruct{}
	err = ctx.BindJSON(changePasswordStruct)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if !services.CheckPasswordHash(changePasswordStruct.PasswordOld, account.Password) {
		ctx.AbortWithError(400, errors.New("PASSWORD_OLD_WRONG"))
		return
	}
	if changePasswordStruct.PasswordNew != changePasswordStruct.PasswordConfirm {
		ctx.AbortWithError(400, errors.New("PASSWORD_NEW_AND_PASSWORD_CONFIRM_NOT_SAME"))
		return
	}

	passwordNewHashed, err := services.HashPassword(changePasswordStruct.PasswordNew)
	if err != nil {
		return
	}
	err = services.UpdateAccount(db, accountID.(int), &models.Account{
		Password: passwordNewHashed,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// UpdateAccount ...
func UpdateAccount(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	account := &models.Account{}
	err := ctx.BindJSON(account)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = services.UpdateAccount(db, accountID.(int), &models.Account{
		Email:   account.Email,
		Name:    account.Name,
		Phone:   account.Phone,
		Address: account.Address,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// UpdateAvatar ...
func UpdateAvatar(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	path, _, err := helpers.UploadFile(ctx)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = services.UpdateAccount(db, accountID.(int), &models.Account{
		Avatar: path,
	})
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.Status(200)
	return
}

// GetPublicAccount ...
func GetPublicAccount(db *gorm.DB, ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	account, err := services.GetAccount(db, int(id))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, models.Account{
		Name:      account.Name,
		Avatar:    account.Avatar,
		Address:   account.Address,
		Phone:     account.Phone,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	})
	return
}

// GetAccount ...
func GetAccount(db *gorm.DB, ctx *gin.Context) {
	accountID, exists := ctx.Get("accountID")
	if !exists {
		ctx.AbortWithError(401, errors.New("Unauthorized"))
		return
	}

	account, err := services.GetAccount(db, accountID.(int))
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	ctx.JSON(200, models.Account{
		ID:        account.ID,
		Email:     account.Email,
		Name:      account.Name,
		Phone:     account.Phone,
		Address:   account.Address,
		Avatar:    account.Avatar,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	})
	return
}
