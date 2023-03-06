package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
)

const createTableStatement = "CREATE TABLE TEMP_TABLE ( NAME VARCHAR2(100), CREATION_TIME TIMESTAMP DEFAULT SYSTIMESTAMP, VALUE  NUMBER(5))"

func main() {
	fmt.Println("Hello, World!")

	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createTableStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Connected to Oracle Database")
}
