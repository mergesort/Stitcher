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

func StitchImages(images []image.Image) {
	imageCount := len(images)

	imageWidth := stitchedImageWidth / imageCount
	imageHeight := stitchedImageHeight / imageCount

	newImage := imaging.New(stitchedImageWidth, stitchedImageWidth, color.NRGBA{0, 0, 0, 255})

	for i := 0; i < imageCount; i++ {
		currImage := images[i]

		x := i / 2 * imageWidth
		y := imageHeight
		if i%2 == 0 {
			y = 0
		}

		croppedImage := CropCenter(currImage)
		resizedImage := imaging.Resize(croppedImage, imageWidth, imageHeight, imaging.BSpline)
		imaging.Overlay(newImage, resizedImage, image.Pt(x, y), 1.0)
	}

	imaging.Save(newImage, "StitchedImage.png")
}
