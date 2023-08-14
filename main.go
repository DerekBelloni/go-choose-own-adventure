package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/derekbelloni/go-choose-own-adventure/handlers"
)


func main() {
	mux := http.NewServeMux()

	storyData := unmarshalJson() 

	mux.Handle("/", &handlers.AdventureHandler{StoryData: storyData})

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