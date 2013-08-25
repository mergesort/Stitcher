package Stitcher

import (
	"image"
	"os"
	"path/filepath"
)

func DownloadImage(url string) image.Image {
	walkDirectory("./")
	fileName := "./tmp/file.zip"
	fmt.Println("Downloading file...")

	output, err := os.Create(fileName)
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)

	fmt.Println(n, "bytes downloaded")
}

func walkDirectory(directory string) {
	fileCount := 0
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) (int, error) {
		fileCount++
		fmt.Println(path, fileCount)
		if err != nil {
			fmt.Println("An error occurred walking the path.", err)
			return 0, err
		}
		return fileCount, nil
	})
}
