package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func openConn() *sql.DB {
	user, _ := os.LookupEnv("MYSQL_USER")
	password, _ := os.LookupEnv("MYSQL_PASSWORD")
	db, _ := os.LookupEnv("MYSQL_DATABASE")
	host, _ := os.LookupEnv("MYSQ_HOST")

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", user, password, host, db))
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(time.Second * 10)
	return conn
}

func InsertUser(conn *sql.DB) {
	db := conn
	username := "johndoe"
	password := "secret"
	createdAt := time.Now()
	_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
}

func selectUser(conn *sql.DB) {

	var (
		id       int =
		username string
		password string
	)

	db := conn
	query := `SELECT id, username, password FROM users WHERE id = ?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn := openConn()
	InsertUser(conn)
	selectUser(conn)

}
