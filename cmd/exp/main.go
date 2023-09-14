package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
	Meta struct {
		Visits int
	}
}

type UserMeta struct {
	Visits int
}

func main() {
	// ParseFiles only sees current directory
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Jojhn",
		Age:  12,
		Bio:  `<script>alert(hello)</script>`,
		Meta: UserMeta{
			Visits: 3,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
