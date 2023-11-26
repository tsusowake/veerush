package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"

	"github.com/tsusowake/veerush/internal/config"
	"github.com/tsusowake/veerush/internal/process"
)

func main() {
	var conf config.Config
	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}

	mainCtx := context.Background()

	dbPoolCtx, cancelDbPoolCtx := context.WithCancel(mainCtx)
	defer cancelDbPoolCtx()

	dbpool, err := pgxpool.New(dbPoolCtx, conf.DB.ConnString())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer dbpool.Close()

	p := process.NewProcess(dbpool)
	if _, err := p.DoVeerush(context.Background()); err != nil {
		panic(err)
	}
}
