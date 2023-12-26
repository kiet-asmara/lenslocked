package controllers

import (
	"html/template"
	"net/http"
)

// closure to create handler
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
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
