package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Indexs() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Join("views", "index.html")
		tmpl, _ := template.ParseFiles(fp)
		tmpl.Execute(w, nil)
	}
}
