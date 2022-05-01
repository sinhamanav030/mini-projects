package story

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("storyTemp.html").ParseFiles("./templates/storyTemp.html"))
}

type Story map[string]Chapter

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}
func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := h.pathFn(req)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter Not Found", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
