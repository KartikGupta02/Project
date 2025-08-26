package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "K@rtik3275"
	dbname   = "db_1"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// Insert a record with hardcoded values
	insertStmt := `INSERT INTO Students (Name, Roll_Number) VALUES ('Jacob', 20)`
	_, err = db.Exec(insertStmt)
	CheckError(err)
	fmt.Println("Inserted hardcoded record")

	// Insert a record dynamically using parameters for safety
	insertDynStmt := `INSERT INTO Students (Name, Roll_Number) VALUES ($1, $2)`
	_, err = db.Exec(insertDynStmt, "Jack", 21)
	CheckError(err)
	fmt.Println("Inserted dynamic record")

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
