package migrations

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func UsersMigrate(db *sql.DB) {
	// users テーブルが存在するかどうかを調べる
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'users'").Scan(&count)
	if err != nil {
		panic(err)
	}

	// users テーブルが存在する場合は処理終了
	if count >= 1 {
		return
	}

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE users (user_id serial PRIMARY KEY, user_name VARCHAR(50));")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		INSERT INTO
			users (user_name)
		VALUES
			('Smith'),
			('ohnson'),
			('Brown')
	`)

	if err != nil {
		panic(err)
	}
}
