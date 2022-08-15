package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("BD")
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	URLTable := `CREATE TABLE IF NOT EXISTS imageboard_db (
		id SERIAL,
		post TEXT,
		unixtime INTEGER);`
	Query, err := db.Prepare(URLTable)
	if err != nil {
		panic(err)
	}
	Query.Exec()
	return db, nil
}
