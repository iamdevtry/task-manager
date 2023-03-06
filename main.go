package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
)

func main() {
	fmt.Println("Hello, World!")
	connectionString := "oracle://" + "system" + ":" + "123456@@" + "@" + "localhost" + ":" + "1521" + "/" + "orcl"
	db, err := sql.Open("oracle", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to Oracle Database")
}
