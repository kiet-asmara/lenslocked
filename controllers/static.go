package controllers

import (
	"net/http"

	"github.com/kiet-asmara/lenslocked/views"
)

// closure to create handler
func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
