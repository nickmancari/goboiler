package main

import (
	"net/http"
	"fmt"
	"html/template"
)

var tpl *template.Template

//Wildcard parse for html files in the directory
func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	//Bring CSS folder into scope
	views := http.FileServer(http.Dir("views"))

	//URL handling
	http.Handle("/views/", http.StripPrefix("/views/", views))

	//URL with func handling (root webpage)
	http.HandleFunc("/", index)

	//Listen and Serve
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWrite, r *http.Request) {
	//Quick serve homepage
	tpl.ExecuteTemplate(w, "index.html", nil)
}
