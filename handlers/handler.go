package handlers

import (
	"net/http"
	"text/template"
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
	ArcName string
	StoryData map[string]Arc
	Template *template.Template
}


func (ah AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	arc, found := ah.StoryData[ah.ArcName]
	if !found {
		http.NotFound(w, r)
		return
	}

	err := ah.Template.Execute(w, arc)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}