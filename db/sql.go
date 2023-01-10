package db

import "database/sql"

func GetDb() *sql.DB {
	file := "/Users/kennethlandsbaum/testing.db"
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}
	return db
}
