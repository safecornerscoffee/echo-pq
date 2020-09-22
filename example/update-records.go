package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	psqlURI := fmt.Sprintf("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	dropTableStatement := `DROP TABLE IF EXISTS users`
	createTableStatement := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			age INT,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE NOT NULL
		);`

	_, err = db.Exec(dropTableStatement)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(createTableStatement)
	if err != nil {
		panic(err)
	}

	createUserStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	returningId := 0
	err = db.QueryRow(createUserStatement, 30, "jon@calhoun.io", "jonathan", "calhoun").Scan(&returningId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Returning id:", returningId)

	updateUserNameByIdStatement := `
	UPDATE users
	SET first_name = $2, last_name = $3
	WHERE id = $1`
	res, err := db.Exec(updateUserNameByIdStatement,
		1, "NewFirst", "NewLast")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated row(s): ", count)

	updateUserNameByIdQueryRowStatement := `
	UPDATE users
	SET first_name = $2, last_name = $3
	WHERE id = $1
	RETURNING id, email;`
	var email string
	var id int
	err = db.QueryRow(updateUserNameByIdQueryRowStatement, returningId, "AnotherFirst", "AnotherLast").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated id=%d, email=%s\n", id, email)

	DeleteUserByIdStatement := `
	DELETE FROM users
	WHERE id = $1`
	res, err = db.Exec(DeleteUserByIdStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted row(s): ", count)
}
