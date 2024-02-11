package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"project/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	fmt.Println("Server is running on port 8080...")
	return http.ListenAndServe(":8080", nil)
}
