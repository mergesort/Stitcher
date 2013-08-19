package Stitcher

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

const (
	stitchedImageHeight = 320
	stitchedImageWidth  = 640
)

func StitchImages(images []image.Image) image.Image {
	imageCount := len(images)

	imageWidth := stitchedImageWidth / imageCount * 2
	imageHeight := stitchedImageHeight / 2

	var newImage *image.NRGBA
	newImage = imaging.New(stitchedImageWidth, stitchedImageHeight, color.NRGBA{255, 0, 0, 255})

	for i := 0; i < imageCount; i++ {
		currImage := images[i]

		x := ((i % (imageCount / 2)) * imageWidth)
		y := 0
		if i >= imageCount/2 {
			y = imageHeight
		}

		croppedImage := CropCenter(currImage)
		resizedImage := imaging.Resize(croppedImage, imageWidth, imageHeight, imaging.CatmullRom)
		newImage = imaging.Paste(newImage, resizedImage, image.Pt(x, y))
	}

	return newImage
}
