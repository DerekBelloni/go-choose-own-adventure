package handlers

import (
	// "fmt"
	"fmt"
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

// I think I might need a handler for each arc
// so maybe something like type IntroHandler struct, maybe it can be handed StoryData['intro']
type AdventureHandler struct{
	StoryData map[string]Arc
	Template *template.Template
}

// type StoryArc struct{
// 	ArcName string
// 	Title string
// 	Body []string
// 	Options []Option
// }

func (ah AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	arcName := "intro"

	// arcName := r.URL.Path

	fmt.Println("request: ", r.URL.Path)

	arc, found := ah.StoryData[arcName]
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