package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	psqlHost := os.Getenv("POSTGRES_HOST")
	psqlPort := os.Getenv("POSTGRES_PORT")
	psqlPassword := os.Getenv("POSTGRES_PASSWORD")
	psqlUser := os.Getenv("POSTGRES_USER")
	psqlDB := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		psqlHost, psqlPort, psqlUser, psqlPassword, psqlDB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 60; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connection Succeed.")
}
