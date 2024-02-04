package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	//　追加
	"example.com/m/api"
	//　追加
	"example.com/m/migrations"
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

	//　追加
	migrations.UsersMigrate(db)
	//　追加
	api.Users(db)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
