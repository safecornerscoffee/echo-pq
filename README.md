# PQL

## Echo

### [database/sql](https://golang.org/pkg/database/sql/#Result)

- sql.Open()
- sql.DB.Exec()
- [sql.Result](https://golang.org/pkg/database/sql/#Result)
- sql.Result.RowsAffected()
- sql.DB.QueryRow()
- [sql.Row](https://golang.org/pkg/database/sql/#Row)
- sql.Row.Scan()
- sql.DB.Query()
- [sql.Rows](https://golang.org/pkg/database/sql/#Rows)

## [Postgres](https://hub.docker.com/_/postgres)

### Environment Variables

```bash
POSTGRES_PASSWORD
POSTGRES_USER
POSTGRES_DB
```

### Initialization scripts

`/docker-entrypoint-initdb.d`

## Refs

- [database/sql documentation](https://golang.org/pkg/database/sql/)
- [lib/pq documentation](https://godoc.org/github.com/lib/pq) 
- [Illustrated guide to SQLX](http://jmoiron.github.io/sqlx/)
- [SQLX - Built in Interfaces](http://jmoiron.net/blog/built-in-interfaces/)
- [Go database/sql tutorial](http://go-database-sql.org/index.html)
- [Using PostgreSQL with Go](https://www.calhoun.io/using-postgresql-with-go)
- [Creating a Go(lang) API with Echo Framework and PostgreSQL](https://www.restapiexample.com/golang-tutorial/creating-golang-api-echo-framework-postgresql)