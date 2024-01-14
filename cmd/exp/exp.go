package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/kiet-asmara/lenslocked/models"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// urutan host user pass gabole kebalik (SASL error)
func (cfg DbConfig) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
}

func main() {
	cfg := DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
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
