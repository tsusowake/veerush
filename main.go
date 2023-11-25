package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"

	"github.com/tsusowake/veerush/internal/config"
	"github.com/tsusowake/veerush/internal/database/generated"
)

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

	queries := generated.New(dbpool)

	u, e := queries.CreateUser(context.Background())
	fmt.Printf("u: %#v\n", u)
	fmt.Printf("e: %#v\n", e)

	uu, ee := queries.GetUserByID(context.Background(), 1)
	fmt.Printf("uu: %#v\n", uu)
	fmt.Printf("ee: %#v\n", ee)
}
