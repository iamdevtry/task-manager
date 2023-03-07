package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iamdevtry/task-manager/api"
	config "github.com/iamdevtry/task-manager/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:  ", err)
	}

	connectionString := fmt.Sprintf("oracle://%s:%s@%s/%s", config.DBUsername, config.DBPassword, config.DBServer, config.DBService)

	db, err := sql.Open(config.DBDriver, connectionString)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := sqlx.NewDb(db, config.DBDriver)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	defer db.Close()
}
