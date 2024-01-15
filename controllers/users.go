package controllers

import (
	"fmt"
	"net/http"

	"github.com/kiet-asmara/lenslocked/models"
)

// models should be where sql lies
type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

// users controller
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")

	// we need a view to render
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}

// sign in controller
func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	// we need a view to render
	u.Templates.SignIn.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	// must be placed before anything is written to writer
	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User authenticated %v", user)
	// no view/template to render, only processing
}

// take user info and print
func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "The email cookie could not be read.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Email cookie: %s\n", cookie.Value) // email
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}
