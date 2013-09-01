package Stitcher

import (
	"encoding/json"
	"fmt"
	"github.com/mergesort/Response"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"sync"
)

///////////////////////////////////////////////////////////////////////////////
// Structs

type IncomingRequest struct {
	ImageURLs []string `json:"urls"`
	Rows      int      `json:"rows"`
	Columns   int      `json:"columns"`
}

///////////////////////////////////////////////////////////////////////////////
// Public functions

func StitcherHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		rw.Header().Set("Content-Type", "application/json")

		if err != nil {
			response := responseJSON.Response{"status": "400", "description": fmt.Sprintf("Your request could not be processed. %v", err), "url": ""}
			sendResponse(rw, response)
		} else {
			var unmarshalledRequest IncomingRequest

			unmarshallingError := json.Unmarshal(body, &unmarshalledRequest)

			if unmarshallingError != nil {
				response := responseJSON.Response{"status": "400", "description": "The data request was formatted incorrectly.", "url": ""}
				sendResponse(rw, response)
			} else {
				for i := 0; i < len(unmarshalledRequest.ImageURLs); i++ {
					currURL := unmarshalledRequest.ImageURLs[i]
					fmt.Println(i, currURL)
				}

				downloadImages(unmarshalledRequest.ImageURLs)
			}
		}
	} else {
		response := responseJSON.Response{"status": "405", "description": "This endpoint requires a POST operation.", "url": ""}
		sendResponse(rw, response)
	}
}

///////////////////////////////////////////////////////////////////////////////
////// Private helpers

func sendResponse(rw http.ResponseWriter, response responseJSON.Response) {
	fmt.Fprint(rw, response)
}

func sendImageToS3(image image.Image) string {
	return ""
}

func downloadImages(imageURLs []string) {
	var group sync.WaitGroup
	images := make([]image.Image, len(imageURLs))

	for _, imageURL := range imageURLs {
		group.Add(1)
		go func() {
			defer group.Done()
			downloadImage(imageURL)
		}()
	}
	group.Wait()

	resultingImage := StitchImages(images)
	sendImageToS3(resultingImage)

	//create slice of images
	//download each image
	//when each one is there, stitch them, then call sendImageToS3
	// DownloadImage(url)
	//when you get a response from s3, call sendResponse with the URL

}

func downloadImage(imageURL string) image.Image {
	response, fetchError := http.Get(imageURL)
	if fetchError != nil {
		fmt.Println("An error occurred downloading ", imageURL, fetchError)
		return nil
	} else {
		defer response.Body.Close()

		image, decodeError := jpeg.Decode(response.Body)

		if decodeError != nil {
			fmt.Println("An error occurred decoding ", imageURL, decodeError)
			return nil
		} else {
			fmt.Println(response)
			fmt.Println(response.Body)
			return image
		}
	}
}
