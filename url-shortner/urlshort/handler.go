package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	return func(w http.ResponseWriter, req *http.Request) {

		path := req.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, req, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, req)
	}
}

// Handler will parse the provided YAML/JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML/JSON, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// JSON is expected to be in format
// 	[
//  	 {"path":"/urlshort","url":"https://github.com/gophercises/urlshort"},
//    	 {"path":"/urlshort-final","url":"https://github.com/gophercises/urlshort/tree/solution"}
// 	]
// The only errors that can be returned all related to having
// invalid YAML/JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func Handler(data []byte, fallback http.Handler, fileType string) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var pathUrls []pathUrl
	var err error
	if fileType == "json" {
		err = parseJSON(&pathUrls, data)
	} else {
		err = parseYaml(&pathUrls, data)
	}
	// pathUrls, err := parseYaml(yml)

	if err != nil {
		return nil, err
	}

	pathToUrls := buildMap(pathUrls)

	return MapHandler(pathToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

func parseJSON(pathUrls *[]pathUrl, data []byte) error {
	// var pathUrls []pathUrl
	err := json.Unmarshal(data, pathUrls)
	if err != nil {
		return err
	}
	fmt.Println(*pathUrls)
	return nil
}

func parseYaml(pathUrls *[]pathUrl, data []byte) error {
	// var pathUrls []pathUrl

	err := yaml.Unmarshal(data, pathUrls)
	if err != nil {
		return err
	}

	fmt.Println(*pathUrls)
	return nil
}

type pathUrl struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
