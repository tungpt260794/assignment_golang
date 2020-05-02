package helpers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile ...
func UploadFile(ctx *gin.Context) (path string, size int64, err error) {
	file, err := ctx.FormFile("image")
	filename := filepath.Base(file.Filename)
	path = fmt.Sprintf(
		"%s/%s_%d_%s",
		"images",
		time.Now().Format("02_01_2006__15_04_05"),
		uuid.New().ID(),
		filename,
	)
	size = file.Size
	if err = ctx.SaveUploadedFile(file, path); err != nil {
		return
	}

	return
}
