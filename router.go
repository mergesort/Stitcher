package Stitcher

import (
	"encoding/json"
	"fmt"
	"github.com/mergesort/Response"
	"io/ioutil"
	"net/http"
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
				response := responseJSON.Response{"status": "200", "description": "Success!", "url": "http://app.net.com/candy"}
				sendResponse(rw, response)
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
