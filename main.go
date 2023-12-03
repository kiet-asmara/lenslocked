package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/kiet-asmara/lenslocked/controllers"
	"github.com/kiet-asmara/lenslocked/views"
)

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse(filepath)
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }

// func contactHandler(w http.ResponseWriter, r *http.Request) {
// 	tplPath := filepath.Join("templates", "contact.gohtml")
// 	executeTemplate(w, tplPath)
// }

// func faqHandler(w http.ResponseWriter, r *http.Request) {
// 	tplPath := filepath.Join("templates", "faq.gohtml")
// 	executeTemplate(w, tplPath)
// }

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

func main() {
	r := chi.NewRouter()

	// parse template
	tpl, err := views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)

}
