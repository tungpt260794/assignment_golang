package routes

import (
	"assignment/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Create ...
func Create(db *gorm.DB) (g *gin.Engine) {
	g = gin.Default()

	v1 := g.Group("/v1")
	{
		v1.GET("", func(ctx *gin.Context) {
			ctx.String(200, "API v1.0")
		})
		v1.POST("/register", func(ctx *gin.Context) { Register(db, ctx) })
		v1.POST("/login", func(ctx *gin.Context) { Login(db, ctx) })

		account := v1.Group("/account")
		{
			account.Use(middlewares.RequireLogin())

			account.GET("", func(ctx *gin.Context) { GetAccount(db, ctx) })
			account.PUT("", func(ctx *gin.Context) { UpdateAccount(db, ctx) })
			account.PUT("/password", func(ctx *gin.Context) { ChangePassword(db, ctx) })
			account.PUT("/avatar", func(ctx *gin.Context) { UpdateAvatar(db, ctx) })
		}

		gallery := v1.Group("/gallery")
		{
			gallery.Use(middlewares.RequireLogin())

			gallery.POST("", func(ctx *gin.Context) { CreateGallery(db, ctx) })
			gallery.PUT("/:id", func(ctx *gin.Context) { UpdateGallery(db, ctx) })
			gallery.PUT("/:id/publication", func(ctx *gin.Context) { PublicGallery(db, ctx) })
		}

		photo := v1.Group("/photo")
		{
			photo.Use(middlewares.RequireLogin())

			photo.POST("", func(ctx *gin.Context) { UploadPhoto(db, ctx) })
			photo.PUT("/:id", func(ctx *gin.Context) { UpdatePhoto(db, ctx) })
			photo.GET("/:id", func(ctx *gin.Context) { GetPhoto(db, ctx) })
			photo.POST("/:id/reaction", func(ctx *gin.Context) { Like(db, ctx) })
			photo.DELETE("/:id/reaction", func(ctx *gin.Context) { UnLike(db, ctx) })
		}

		public := v1.Group("/public")
		{
			public.GET("/account/:id", func(ctx *gin.Context) { GetPublicAccount(db, ctx) })
			public.GET("/gallery", func(ctx *gin.Context) { GetPublicGalleries(db, ctx) })
			public.GET("/gallery/:id", func(ctx *gin.Context) { GetPublicGallery(db, ctx) })
			public.GET("/photo/:id/w/:width", func(ctx *gin.Context) { DownloadPhoto(db, ctx) })
		}
	}

	return
}
