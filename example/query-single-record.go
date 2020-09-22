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

type User struct {
	ID        string
	Age       string
	FirstName string
	LastName  string
	Email     string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	queryStatement := `
	SELECT id, email from users
	WHERE id=$1;`
	var email string
	var id int
	row := db.QueryRow(queryStatement, 9999)
	err = row.Scan(&id, &email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}
	queryStatement = `
	SELECT * from users
	WHERE id=$1;`
	var user User
	row = db.QueryRow(queryStatement, 3)
	err = row.Scan(&user.ID, &user.Age, &user.FirstName,
		&user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
}
