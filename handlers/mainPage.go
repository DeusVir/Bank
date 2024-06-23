package mainPage

import (
	"html/template"
	"log"
	"net/http"
)


type Start struct{
	l *log.Logger
}

func NewStart(l *log.Logger) *Start{
	return &Start{l}
}

func (h *Start) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, err := template.ParseFiles("D:/Bank/templates/html/index.html")

	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
	}

	err = t.Execute(w, nil)

	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
	}

}