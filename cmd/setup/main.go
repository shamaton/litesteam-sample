package main

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/shamaton/litestream-sample/model"
	"github.com/shamaton/litestream-sample/sqlite"
)

func main() {

	sqlite.RegisterDriver()
	h := dbHandler()
	defer h.Close()

	// create table
	if err := h.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// remove table
}

func dbHandler() *model.Client {
	if err := os.MkdirAll("tmp/sqlite", 0777); err != nil {
		log.Fatalf("failed creating directory: %v", err)
	}

	var options []model.Option
	options = append(options, model.Debug())
	client, err := model.Open(dialect.SQLite, "file:./tmp/sqlite/db.sqlite?cache=shared", options...)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	return client
}
