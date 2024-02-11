package main

import (
	"database/sql"
	"fmt"
	"os"

	"project/api"
	"project/controllers"
	"project/migrations"

	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("POSTGRES_HOSTNAME")
	database := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, database))
	if err != nil {
		panic(err)
	}

	migrations.UsersMigrate(db)
	api.Users(db)

	controllers.StartMainServer()
}
