package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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
func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	tpl := template.New(pattern[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrf field not implemented")
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{htmlTpl: tpl}, nil
}

// func Parse(filepath string) (Template, error) {
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing template: %w", err)
// 	}
// 	return Template{htmlTpl: tpl}, nil
// }

// template type
type Template struct {
	htmlTpl *template.Template
}

// execute parsed template
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.htmlTpl.Clone() //clone erases race condition since htmltpl is a pointer
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}

	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r) // inject function to html
			},
		},
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buff bytes.Buffer // takes a bit of memory, unsuitable for large pages
	err = tpl.Execute(&buff, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buff)
}

// steps in executing a function in template (csrf in this case)
// 1. define a placeholder function in parseFS before template parse
// 2. update placeholder in execute
