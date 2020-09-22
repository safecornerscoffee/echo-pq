package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// type User struct {
// 	ID        string
// 	Age       string
// 	FirstName string
// 	LastName  string
// 	Email     string
// }

func main() {
	psqlURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)

	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, firstName)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
