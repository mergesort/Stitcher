package Stitcher

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
)

const (
	squareSide = 640
)

func CropCenter(sourceImage image.Image) {
	x0 := (sourceImage.Bounds().Size().X - squareSide) / 2
	y0 := (sourceImage.Bounds().Size().Y - squareSide) / 2

	croppedPath := fmt.Sprintf("Cropped-%s", filePath)
	cropTo640Squared(sourceImage, croppedPath, x0, y0)
}

func CropFromTop(sourceImage image.Image) {
	x0 := 0
	y0 := 0

	croppedPath := fmt.Sprintf("Cropped-%s", filePath)
	cropTo640Squared(sourceImage, croppedPath, x0, y0)
}

func CropFromBottom(sourceImage image.Image) {
	x0 := sourceImage.Bounds().Size().X - squareSide
	y0 := sourceImage.Bounds().Size().Y - squareSide

	croppedPath := fmt.Sprintf("Cropped-%s", filePath)
	cropTo640Squared(sourceImage, croppedPath, x0, y0)
}

////////////////////////////////////////////////////////////////////////////////
//Private helpers

func cropImage(sourceImage image.Image, x0, y0, x1, y1 int) image.Image {
	currImage := imaging.Clone(sourceImage)
	croppedImage := imaging.Crop(currImage, image.Rect(x0, y0, x1, y1))
	return croppedImage
}

func cropTo640Squared(sourceImage image.Image, outputFileName string, x0, y0 int) {
	x1 := squareSide + x0
	y1 := squareSide + y0

	finalImage := cropImage(sourceImage, x0, y0, x1, y1)
	imaging.Save(finalImage, outputFileName)
}
