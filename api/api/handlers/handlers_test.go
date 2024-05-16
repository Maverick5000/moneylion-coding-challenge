package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFormatResponse(t *testing.T) {
	incomingPayload := IncomingPayload{
		ContentCards: []struct {
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
		}{
			{
				Id:       "1",
				ImageUri: "https://example.com/image1.jpg",
				TextData: struct {
					Title    string `json:"title"`
					Subtitle string `json:"subTitle"`
					Body     string `json:"body"`
					Author   struct {
						First string `json:"first"`
						Last  string `json:"last"`
					} `json:"author"`
				}{
					Title:    "Sample Title 1",
					Subtitle: "Sample Subtitle 1",
					Body:     "Sample Body 1",
					Author: struct {
						First string `json:"first"`
						Last  string `json:"last"`
					}{
						First: "John",
						Last:  "Doe",
					},
				},
				Metadata: struct {
					PublishDate string `json:"publishDate"`
					Priority    int    `json:"priority"`
				}{
					PublishDate: "2024-05-20",
					Priority:    1,
				},
				Comments: []Comment{
					{
						Text:       "Comment 1",
						Author:     "Jane Doe",
						ProfilePic: "https://example.com/profile1.jpg",
						Likes:      5,
					},
				},
			},
		},
	}

	result := formatResponse(incomingPayload)

	expected := []ContentCard{
		{
			Id:          "1",
			ImageUri:    "https://example.com/image1.jpg",
			Title:       "Sample Title 1",
			Subtitle:    "Sample Subtitle 1",
			Body:        "Sample Body 1",
			Author:      "John Doe",
			PublishDate: "2024-05-20",
			Priority:    1,
			Comments: []Comment{
				{
					Text:       "Comment 1",
					Author:     "Jane Doe",
					ProfilePic: "https://example.com/profile1.jpg",
					Likes:      5,
				},
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("formatResponse() returned unexpected result\nExpected: %v\nGot: %v", expected, result)
	}
}

func TestFeedHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com", nil)
	rr := httptest.NewRecorder()

	FeedHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}

	var cards []ContentCard
	err := json.NewDecoder(rr.Body).Decode(&cards)
	if err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	expectedNumCards := 5
	if len(cards) != expectedNumCards {
		t.Errorf("handler returned unexpected number of cards: got %v want %v",
			len(cards), expectedNumCards)
	}
}
