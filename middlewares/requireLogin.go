package middlewares

import (
	"errors"
	"strings"

	"assignment/services"

	"github.com/gin-gonic/gin"
)

// RequireLogin ...
func RequireLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		arr := strings.Split(token, " ")

		if len(arr) != 2 {
			ctx.AbortWithError(401, errors.New("Unauthorized"))
			return

		}

		if arr[0] != "Bearer" {
			ctx.AbortWithError(401, errors.New("Unauthorized"))
			return
		}
		token = arr[1]
		claims, err := services.VerifyToken(token)

		if err != nil {
			ctx.AbortWithError(401, errors.New("Unauthorized"))
			return
		}
		ctx.Set("accountID", claims.AccountID)

		ctx.Next()
	}
}
