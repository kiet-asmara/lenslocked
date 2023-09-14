package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Parsing template %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("Templates", "home.gohtml")
	executeTemplate(w, tplPath) // generally make execute last action
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("Templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<h1>FAQ Page</h1>
	<ul>
		<li><b>Is there a free version?</b> Yes there is bro</li>
		<li><b>What are your support hours?</b> From 9AM-6PM</li>
	</ul>
	`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("start at :3000")
	http.ListenAndServe(":3000", r)
}

// func main() {
// 	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello"))
// 	}

// 	http.HandleFunc("/", handlerIndex)
// 	address := "localhost:9000"
// 	fmt.Println("start server at ", address)
// 	err := http.ListenAndServe(address, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// }
