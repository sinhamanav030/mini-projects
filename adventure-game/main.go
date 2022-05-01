package main

import (
	"adventure-game/story"
	"flag"
	"fmt"
	// "html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to run application")
	filename := flag.String("file", "./gopher.json", "the json file with story")
	flag.Parse()
	fmt.Println("Using the story from ", *filename)
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	st, err := story.JsonStory(f)
	if err != nil {
		panic(err)
	}
	// tpl := template.Must(template.New("").Parse("Hello"))
	handler := story.NewHandler(st)
	fmt.Println("Starting the server at port :", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
