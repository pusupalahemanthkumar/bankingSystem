package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pusupalahemanthkumar/bankingsystem/util"
)

var testQueries *Queries

var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../../")

	if err != nil {
		fmt.Println("Cannot Load Config Data : ", err)
		log.Fatal("Cannot Load Config Data : ", err)
	}

	testDB, err = pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		fmt.Println("Cannot Connect With DB : ", err)
		log.Fatal("Cannot Connect With DB : ", err)
	}

	testQueries = New(testDB)
	fmt.Print("Testing connection !!")
	os.Exit(m.Run())

}
