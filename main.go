package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tsusowake/veerush/internal/config"
	"log"
)
import "github.com/kelseyhightower/envconfig"

func main() {
	var conf config.Config
	if err := envconfig.Process("", &conf); err != nil {
		log.Fatal(err)
	}

	dbpool, err := pgxpool.New(context.Background(), conf.DB.ConnString())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer dbpool.Close()

	var userID string
	err = dbpool.QueryRow(context.Background(), "select id from users").Scan(&userID)
	if err != nil {
		panic(fmt.Sprintf("QueryRow failed: %v\n", err))
	}

	fmt.Println(userID)
}
