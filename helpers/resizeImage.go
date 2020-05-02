package helpers

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/gift"
	"github.com/google/uuid"
)

// ImageSize ...
type ImageSize struct {
	Width int
	Path  string
}

// Widths ...
var Widths = [...]int{1920, 1600, 1280, 1024, 800, 256}

// ResizeAll ...
func ResizeAll(src string) (sizes map[int]string, err error) {
	sizes = map[int]string{}

	w, _, err := GetDimension(src)
	if err != nil {
		return
	}

	var cimg = make(chan ImageSize)
	counter := 0

	for _, width := range Widths {
		if w > width {
			counter++
			go func(w int, cimg chan ImageSize) {
				dst, err := Resize(src, w, 0)
				if err != nil {
					return
				}
				cimg <- ImageSize{
					Width: w,
					Path:  dst,
				}
			}(width, cimg)
		}
	}

	for ; counter > 0; counter-- {
		is := <-cimg
		sizes[is.Width] = is.Path
	}

	return
}

// GetDimension ...
func GetDimension(src string) (int, int, error) {
	file, err := os.Open(src)
	if err != nil {
		return 0, 0, err
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}

// Resize ...
func Resize(src string, width int, height int) (dst string, err error) {
	g := gift.New(
		gift.Resize(width, height, gift.LanczosResampling),
	)
	imageFile, err := loadImage(src)
	if err != nil {
		return
	}
	imageDst := image.NewRGBA(g.Bounds(imageFile.Bounds()))
	g.Draw(imageDst, imageFile)

	dst = fmt.Sprintf(
		"%s/%s_%s_%d_w%d.jpg",
		filepath.Dir(src),
		filenameWithoutExt(filepath.Base(src)),
		time.Now().Format("02_01_2006__15_04_05"),
		uuid.New().ID(),
		width,
	)
	err = saveImage(dst, imageDst)

	return
}

func loadImage(src string) (img image.Image, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()
	img, _, err = image.Decode(f)
	if err != nil {
		return
	}
	return
}

func saveImage(src string, img image.Image) (err error) {
	f, err := os.Create(src)
	if err != nil {
		return
	}
	defer f.Close()
	err = jpeg.Encode(f, img, &jpeg.Options{
		Quality: 80,
	})
	if err != nil {
		return
	}

	return
}

func filenameWithoutExt(fn string) string {
	return strings.TrimSuffix(fn, filepath.Ext(fn))
}
