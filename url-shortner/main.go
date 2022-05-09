package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"url-shortner/urlshort"
)

var (
	handler http.HandlerFunc
	mux     *http.ServeMux
)

func main() {
	mux = defaultMux()

	// Build the MapHandler using the mux as the fallback
	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	pathsToUrls := getInfo()
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	handler = setHandler(mapHandler)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	fmt.Println("Starting the server on :8080")

	http.HandleFunc("/add", addUrl)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func setHandler(mapHandler http.HandlerFunc) http.HandlerFunc {
	var handler http.HandlerFunc
	if len(os.Args) > 1 {
		file := os.Args[1]
		fs, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err, "please enter file as location")
			return nil
		}
		data := string(fs)

		if strings.Contains(file, ".json") {
			handler, err = urlshort.Handler([]byte(data), mapHandler, "json")
			if err != nil {
				log.Fatal(err)
				return nil
			}
		} else if strings.Contains(file, ".yaml") {
			handler, err = urlshort.Handler([]byte(data), mapHandler, "yaml")
			if err != nil {
				log.Fatal(err)
				return nil
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

	return handler

}

func addUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		url := r.FormValue("url")
		path := r.FormValue("path")
		addInfo(url, path)
		pathsToUrls := getInfo()
		mapHandler := urlshort.MapHandler(pathsToUrls, mux)
		handler = setHandler(mapHandler)
		http.Redirect(w, r, "added info", http.StatusSeeOther)
		return
	}
	tpl.Execute(w, nil)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
