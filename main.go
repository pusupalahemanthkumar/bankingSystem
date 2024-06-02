package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pusupalahemanthkumar/bankingsystem/api"
	db "github.com/pusupalahemanthkumar/bankingsystem/db/sqlc"
	"github.com/pusupalahemanthkumar/bankingsystem/util"
)

func main() {
	fmt.Println("Banking System project")

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config !!", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("connection with postgresdb failed !!", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server !!", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server !!", err)
	}

}
