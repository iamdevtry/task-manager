package main

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/iamdevtry/task-manager/util"
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	fmt.Println("Hello, World!")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf("oracle://%s:%s@%s/%s", config.DBUsername, config.DBPassword, config.DBServer, config.DBService)

	db, err := sql.Open(config.DBDriver, connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Connected to Oracle Database")
}
