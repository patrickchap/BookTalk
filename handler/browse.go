package handler

import (
	"encoding/json"
	"fmt"
	"github/patrickchap/booktalk/model"
	"github/patrickchap/booktalk/view"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BrowseHandler struct{}

func (h *BrowseHandler) Index(c echo.Context) error {
	baseURL := "https://www.googleapis.com/books/v1/volumes?q=subject:fantasy&startIndex=0&maxResults=6"

	// Make the API request
	response, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer response.Body.Close()
	//Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	var booksData map[string]interface{}
	err = json.Unmarshal(body, &booksData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	// Get the items array from the response
	items, ok := booksData["items"].([]interface{})
	if !ok {
		fmt.Println("Items not found in JSON data.")
	}

	var books []model.Book
	for _, item := range items {
		bookData, ok := item.(map[string]interface{})
		if !ok {
			fmt.Println("Book data not found in item.")
			continue
		}
		volumeInfo, ok := bookData["volumeInfo"].(map[string]interface{})
		if !ok {
			fmt.Println("Items not found in JSON data.")
		}

		title, ok := volumeInfo["title"].(string)
		if !ok {
			fmt.Println("Title not found in volumeInfo.")
		}
		authors, ok := volumeInfo["authors"].([]interface{})
		if !ok {
			fmt.Println("Authors not found in volumeInfo.")
		}
		author, ok := authors[0].(string)
		if !ok {
			fmt.Println("Author not found in authors.")
		}

		imageLinks, ok := volumeInfo["imageLinks"].(map[string]interface{})

		if !ok {
			fmt.Println("ImageLinks not found in volumeInfo.")
		}

		thumbnail, ok := imageLinks["thumbnail"].(string)
		if !ok {
			fmt.Println("Thumbnail not found in imageLinks.")
		}

		books = append(books, model.Book{
			Title:     title,
			Author:    author,
			Thumbnail: thumbnail,
		})

	}

	return render(c, view.Browse(books))

}
