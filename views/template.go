package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

// must added to create error that can panic without cluttering code
// and return a template to parse
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// errors can be wrapped to include other messages (strings) inside
// used to hide details/add abstraction or detail the error??

// difference is parse fs to accept variadic parameters to detect embedded commands
// eg: from parent folder run lenslocked/app.exe
// couldnt be read before since relative path
func ParseFS(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

// template type
type Template struct {
	htmlTpl *template.Template
}

// execute parsed template
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}
