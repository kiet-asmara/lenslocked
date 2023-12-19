package controllers

import (
	"html/template"
	"net/http"

	"github.com/kiet-asmara/lenslocked/views"
)

// closure to create handler
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML // not good w/script injection
	}{
		{
			Question: "Is there a free version",
			Answer:   "Yes",
		},
		{
			Question: "Is there a free version",
			Answer:   "Yes",
		},
		{
			Question: "How can we contact?",
			Answer:   `Email us - <a href="mailto:kiet123pascal@gmail.com">kiet123pascal@gmail.com</a>`,
		},
		{
			Question: "How can we contact?",
			Answer:   `Email us - <a href="mailto:kiet123pascal@gmail.com">kiet123pascal@gmail.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
