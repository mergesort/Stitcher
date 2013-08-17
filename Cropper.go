package Stitcher

import (
	"github.com/disintegration/imaging"
	"image"
)

const (
	squareSide = 640
)

func CropCenter(sourceImage image.Image) image.Image {
	x0 := (sourceImage.Bounds().Size().X - squareSide) / 2
	y0 := (sourceImage.Bounds().Size().Y - squareSide) / 2

	croppedImage := cropTo640Squared(sourceImage, x0, y0)
	return croppedImage
}

func CropFromTop(sourceImage image.Image) image.Image {
	x0 := 0
	y0 := 0

	croppedImage := cropTo640Squared(sourceImage, x0, y0)
	return croppedImage
}

func CropFromBottom(sourceImage image.Image) image.Image {
	x0 := sourceImage.Bounds().Size().X - squareSide
	y0 := sourceImage.Bounds().Size().Y - squareSide

	croppedImage := cropTo640Squared(sourceImage, x0, y0)
	return croppedImage
}

////////////////////////////////////////////////////////////////////////////////
//Private helpers

func cropImage(sourceImage image.Image, x0, y0, x1, y1 int) image.Image {
	currImage := imaging.Clone(sourceImage)
	croppedImage := imaging.Crop(currImage, image.Rect(x0, y0, x1, y1))

	return croppedImage
}

func cropTo640Squared(sourceImage image.Image, x0, y0 int) image.Image {
	x1 := squareSide + x0
	y1 := squareSide + y0

	finalImage := cropImage(sourceImage, x0, y0, x1, y1)
	return finalImage
}
