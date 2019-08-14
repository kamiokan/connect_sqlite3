package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// テーブル作成
// CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))
func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Pingの送信
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// データの挿入
	res, err := db.Exec(
		`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		123,
		"title",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
