package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/shamaton/litestream-sample/db"
	"github.com/shamaton/litestream-sample/sqlite"
)

const dbFilePath = "/data/db.sqlite"

var (
	addr = ":3000"
)

func main() {

	// whether to create tables
	shouldCreateTable := false
	if _, err := os.Stat(dbFilePath); err != nil {
		shouldCreateTable = true
	}

	// prepare db handler
	sqlite.RegisterDriver()
	dbh, err := db.Open(dialect.SQLite, "file:"+dbFilePath+"?cache=shared", db.Debug())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// create table
	if shouldCreateTable {
		if err := dbh.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		log.Println("created tables.")
	}

	// prepare server
	h := handler{dbh: dbh}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/insert", h.InsertRecord)
	r.Get("/select", h.SelectRecord)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// start server
	log.Println("server start", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

type handler struct {
	dbh *db.Client
}

func (h handler) InsertRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := h.dbh.User.
		Create().
		SetAge(int(time.Now().UnixNano() % 100)).
		SetName(fmt.Sprintf("at%d", time.Now().Unix())).
		Save(ctx)

	if err != nil {
		writeError(w, fmt.Sprintf("failed creating user: %v", err))
		return
	}

	writeJSON(w, user)
}

func (h handler) SelectRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.dbh.User.Query().All(ctx)
	if err != nil {
		writeError(w, fmt.Sprintf("failed selecting users: %v", err))
		return
	}
	writeJSON(w, users)
}

func writeJSON(w http.ResponseWriter, v any) {

	b, err := json.Marshal(v)
	if err != nil {
		writeError(w, fmt.Sprintf("failed unmarshaling: %v", err))
		return
	}

	writeResponse(w, http.StatusOK, b)
}

func writeError(w http.ResponseWriter, message string) {
	writeResponse(w, http.StatusInternalServerError, []byte(message))
}

func writeResponse(w http.ResponseWriter, code int, b []byte) {
	w.WriteHeader(code)
	if _, err := w.Write(b); err != nil {
		log.Printf("failed writing response: %v\n", err)
	}
}
