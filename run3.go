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

	// 複数行をSELECTする
	rows, err := db.Query(`SELECT id, name, age FROM users ORDER BY name`)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int64
		var name string
		var age int64
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, age)
	}

	// 単一行をSELECTする
	row := db.QueryRow(`SELECT name, age FROM users WHERE id=$1`, 3)
	var name string
	var age int64
	err = row.Scan(&name, &age)
	fmt.Println(name, age)
}
