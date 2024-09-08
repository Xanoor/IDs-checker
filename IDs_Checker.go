package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ///////////////////////////////
// YOU HAVE TO MODIFY THIS PART !
var start int = 0                                 //Start scan from 0
var end int = 100                                 // Check up to ID 100
var errorMsg string = "Post not found"            // Message for error
var url string = "https://example.com/myPage&id=" // URL without IDs to test
var cookieValue string = "session=XXXXXXXXXXXXXX" // If you want to add a cookie (for connection...)

/////////////////////////////////

var ids []int // Store all valid IDs - don't touch this

func main() {
	actualId := start

	for actualId <= end {
		fullUrl := fmt.Sprintf("%s%d", url, actualId)

		req, err := http.NewRequest("GET", fullUrl, nil)
		if err != nil {
			panic(err) // Error during creation of request
		}

		req.Header.Add("Cookie", cookieValue)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err) // Error during request
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic("Error during execution!")
			}

			if !strings.Contains(string(body), errorMsg) {
				ids = append(ids, actualId)
			}
		}
		fmt.Printf("\rPerforming request for ID %d", actualId)
		actualId++ // Next ID
	}

	fmt.Println("\nList of valid IDs:", ids)
}
