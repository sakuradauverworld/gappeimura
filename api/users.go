package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type User struct {
	Id   int    `db:"user_id"`
	Name string `db:"user_name"`
}

func Users(db *sql.DB) {
	http.HandleFunc("/users/get", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			user := User{}
			err := rows.Scan(&user.Id, &user.Name)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	})

	http.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		user := User{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Fatal(err)
		}

		stmt, err := db.Prepare("INSERT INTO users (user_name) VALUES ($1) RETURNING user_id, user_name;")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		err = stmt.QueryRow(user.Name).Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(user)
	})
}
