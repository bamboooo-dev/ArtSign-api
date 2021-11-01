package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"artsign"
	"artsign/app/interface/mysql"
	"artsign/pkg/env"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

var revision = "undefined"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic '%v' captured\n", err)
		}
	}()

	fmt.Printf("Version is %s\n", revision)

	cfg, err := env.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error(err.Error())
			return
		}
	}()

	sugar := logger.Sugar()

	artsignDB, err := mysql.NewDB(cfg.ArtsignMySQL)
	if err != nil {
		sugar.Error(ctx, err)
		return
	}
	entClient, err := mysql.NewClient(artsignDB)
	if err != nil {
		sugar.Error(ctx, err)
		return
	}
	defer func() {
		if err := entClient.Close(); err != nil {
			sugar.Error(ctx, err)
			return
		}
	}()

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(artsign.NewSchema(entClient))
	srv.Use(entgql.Transactioner{TxOpener: entClient})
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
