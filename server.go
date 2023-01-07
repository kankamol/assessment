package main

import (
	"fmt"
	"os"
	"database/sql"

)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
	db, err := sql.Open("postgres",os.Getenv(DATABASE_URL))
	if err != nil {
		log.Fatal("Connect to Database Error", err)
	}
	defer db.Close()

	log.Println("okay")

	createTb := `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);`
	_, err =db.Exec(createTb)
	if err != nil {
		log.Fatal("cant create table", err)
	}
	fmt.Println("create table success")
}
