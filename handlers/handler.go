package handlers

import (
	"fmt"
	"net/http"
)

type Arc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options,omitempty"`
}

type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type AdventureHandler struct{
	StoryData map[string]Arc
}

func (ah AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Look at the request to see what part of the adventure the user has chosen and grab the appropriate data
	// Hand back the template as well
	for key, value := range ah.StoryData {
		fmt.Printf("Story key: %s\nTitle: %s\nStory: %v\nOptions: %v\n", key, value.Title, value.Story, value.Options)
	}

	fmt.Println("Hello World!")
}