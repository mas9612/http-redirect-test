package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/301", moved)
	http.HandleFunc("/302", found)
	http.HandleFunc("/303", seeOther)
	http.HandleFunc("/307", temporary)
	http.HandleFunc("/308", permanent)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("/\tMethod %s\n\n", req.Method)

	data := struct {
		Method string
	}{
		req.Method,
	}

	tpl.ExecuteTemplate(w, "index.html", data)
}

func moved(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("/301\tMethod %s\n", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func found(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("/302\tMethod %s\n", req.Method)
	http.Redirect(w, req, "/", http.StatusFound)
}

func seeOther(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("/303\tMethod %s\n", req.Method)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func temporary(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("/307\tMethod %s\n", req.Method)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func permanent(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("/308\tMethod %s\n", req.Method)
	http.Redirect(w, req, "/", http.StatusPermanentRedirect)
}
