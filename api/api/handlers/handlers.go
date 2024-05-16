package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type Comment struct {
	Text       string `json:"text"`
	Author     string `json:"author"`
	ProfilePic string `json:"profilePic"`
	Likes      int    `json:"likes"`
}

type ContentCard struct {
	Id          string    `json:"id"`
	ImageUri    string    `json:"imageUri"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Body        string    `json:"body"`
	Author      string    `json:"author"`
	PublishDate string    `json:"publishDate"`
	Comments    []Comment `json:"comments"`
	Priority    int       `json:"priority"`
}

type IncomingPayload struct {
	ContentCards []struct {
		Id       string `json:"id"`
		ImageUri string `json:"imageUri"`
		TextData struct {
			Title    string `json:"title"`
			Subtitle string `json:"subTitle"`
			Body     string `json:"body"`
			Author   struct {
				First string `json:"first"`
				Last  string `json:"last"`
			} `json:"author"`
		} `json:"textData"`
		Metadata struct {
			PublishDate string `json:"publishDate"`
			Priority    int    `json:"priority"`
		} `json:"metadata"`
		Comments []Comment `json:"comments"`
	} `json:"contentCards"`
}

func formatResponse(incomingPayload IncomingPayload) []ContentCard {
	var data []ContentCard
	for _, item := range incomingPayload.ContentCards {
		card := ContentCard{
			Id:          item.Id,
			ImageUri:    item.ImageUri,
			Title:       item.TextData.Title,
			Subtitle:    item.TextData.Subtitle,
			Body:        item.TextData.Body,
			Author:      fmt.Sprintf("%s %s", item.TextData.Author.First, item.TextData.Author.Last),
			PublishDate: item.Metadata.PublishDate,
			Priority:    item.Metadata.Priority,
			Comments:    item.Comments,
		}
		data = append(data, card)
	}

	// order by priority
	sort.Slice(data, func(i, j int) bool {
		return data[i].Priority > data[j].Priority
	})

	return data
}

func FeedHandler(w http.ResponseWriter, r *http.Request) {
	// Make an HTTP GET request to the specified URL
	resp, err := http.Get("https://stoplight.io/mocks/engine/fullstack-spec/52502230/content")
	if err != nil {
		// If there is an error in the HTTP request, return a 500 Internal Server Error response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Ensure that the response body is closed when the function returns
	defer resp.Body.Close()

	// Define a variable to hold the incoming payload data
	var incomingPayload IncomingPayload
	// Decode the JSON response body into the incomingPayload variable
	err = json.NewDecoder(resp.Body).Decode(&incomingPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format the incoming payload data into the desired response format
	data := formatResponse(incomingPayload)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Encode the data into JSON and write it to the response body
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
