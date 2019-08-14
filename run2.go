package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// テーブル作成
// DROP TABLE IF EXISTS users;
// CREATE TABLE IF NOT EXISTS users (
// id INTEGER PRIMARY KEY,
// name TEXT NOT NULL,
// age INTEGER NOT NULL);
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
		`INSERT INTO users (name, age) VALUES (?, ?)`,
		"Bob",
		18,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 更新件数をチェック
	affected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("更新件数:", affected)

	// 最終挿入ID
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("最終挿入ID:", lastInsertID)
}
