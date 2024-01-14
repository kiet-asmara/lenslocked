package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/kiet-asmara/lenslocked/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg) // postgres code is now separated
	if err != nil {             // easier for ppl, no install driver
		panic(err)
	}
	defer db.Close()

	// test request to db w/ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to db")

	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("john@gmail.com", "12345")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
