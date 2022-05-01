package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"url-shortner/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	var handler http.HandlerFunc

	if len(os.Args) > 1 {
		file := os.Args[1]
		fs, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err, "please enter file as location")
			return
		}
		data := string(fs)

		if strings.Contains(file, ".json") {
			handler, err = urlshort.Handler([]byte(data), mapHandler, "json")
			if err != nil {
				log.Fatal(err)
				return
			}
		} else if strings.Contains(file, ".yaml") {
			handler, err = urlshort.Handler([]byte(data), mapHandler, "yaml")
			if err != nil {
				log.Fatal(err)
				return
			}
		} else {
			fmt.Println("Cannot handle file", file)
			fmt.Println("Starting server on default configuration")
			handler = mapHandler
		}
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Starting server on default configuration")
		handler = mapHandler
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
