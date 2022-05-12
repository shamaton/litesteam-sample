package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"github.com/shamaton/litestream-sample/model"
)

func main() {
	registerSQLite()
	client := dbHandler()

	cmd := os.Getenv("CMD")
	switch cmd {
	case "create":
		createTable(client)

	case "insert":
		insertRecord(client)

	case "select":
		selectRecord(client)
	}
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

func createTable(client *model.Client) {

	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func insertRecord(client *model.Client) {
	ctx := context.Background()
	user, err := client.User.
		Create().
		SetAge(20).
		SetName(fmt.Sprintf("at%d", time.Now().Unix())).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Println("user was created: ", user)
}

func selectRecord(client *model.Client) {
	ctx := context.Background()
	users, err := client.User.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed selecting users: %v", err)
	}
	for _, user := range users {
		log.Println("user :", user.String())
	}
}
