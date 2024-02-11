package controllers

import (
	"html/template"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/templates/top.html")
	t.Execute(w, nil)
}
