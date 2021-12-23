package cyoa

import (
	"html/template"
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

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template.html"))

	path := r.URL.Path

	if path == "/" || path == "" {
		path = "/intro"
	}
	path = path[1:]

	chapter := h.s[path]
	tmpl.Execute(w, chapter)
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}
