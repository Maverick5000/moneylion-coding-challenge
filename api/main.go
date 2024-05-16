package main

import (
	"fmt"
	"moneylion-conding-challenge/api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Registering the FeedHandler function to handle GET requests to /feed route
	mux.HandleFunc("GET /feed", handlers.FeedHandler)

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", mux); err != nil {
		// Print any error that occurs during server startup
		fmt.Println(err.Error())
	}
}

/*
The route /feed returns an array of JSON objects, each representing a content card with the following structure:

{
  "id": "string",
  "imageUri": "string",
  "title": "string",
  "subtitle": "string",
  "body": "string",
  "author": "string",
  "publishDate": "string",
  "comments": [
    {
      "text": "string",
      "author": "string",
      "profilePic": "string",
      "likes": 0
    }
  ],
  "priority": 0
}

Where:
- id: The unique identifier of the content card.
- imageUri: The URI of the image associated with the content card.
- title: The title of the content card.
- subtitle: The subtitle or secondary information of the content card.
- body: The main body content of the content card.
- author: The author of the content card.
- publishDate: The date when the content card was published.
- comments: An array of JSON objects representing comments on the content card, each with text, author, profilePic, and likes fields.
- priority: The priority level of the content card.
*/
