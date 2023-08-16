package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/derekbelloni/go-choose-own-adventure/handlers"
)


func main() {
	mux := http.NewServeMux()

	storyData := unmarshalJson() 

	tmpl, err := template.ParseFiles("./templates/arc.page.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", &handlers.AdventureHandler{
		ArcName: "intro",
		StoryData: storyData,
		Template: tmpl,
	})
	mux.Handle("/new-york", &handlers.AdventureHandler{
		ArcName: "new-york",
		StoryData: storyData,
		Template: tmpl,
	})
	mux.Handle("/debate", &handlers.AdventureHandler{
		ArcName: "debate",
		StoryData: storyData,
		Template: tmpl,
	})
	mux.Handle("/sean-kelly", &handlers.AdventureHandler{
		ArcName: "sean-kelly",
		StoryData: storyData,
		Template: tmpl,
	})
	mux.Handle("/mark-bates", &handlers.AdventureHandler{
		ArcName: "mark-bates",
		StoryData: storyData,
		Template: tmpl,
	})
	mux.Handle("/denver", &handlers.AdventureHandler{
		ArcName: "denver",
		StoryData: storyData,
		Template: tmpl,
	})
	mux.Handle("/home", &handlers.AdventureHandler{
		ArcName: "home",
		StoryData: storyData,
		Template: tmpl,
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func unmarshalJson() map[string]handlers.Arc {
	jsonFile, err := os.ReadFile("adventure.json")
	if err != nil {
		log.Fatal(err)
	}

	var storyData map[string]handlers.Arc

	err = json.Unmarshal([]byte(jsonFile), &storyData)

	if err != nil {
		fmt.Println("error: ", err)
	}

	return storyData
}