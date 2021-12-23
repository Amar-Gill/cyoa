package cyoa

import (
	"html/template"
	"log"
	"net/http"
)

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}

type Story map[string]Chapter

type handler struct {
	s Story
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("template.html"))
}

// assign ServeHTTP method to handler struct, so that it implements http.Handler interface
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" || path == "" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}
