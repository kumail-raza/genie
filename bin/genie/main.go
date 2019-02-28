package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mohuk/genie/formly"

	"github.com/mohuk/genie/handlers"
	"github.com/mohuk/genie/manager"

	"github.com/gorilla/mux"

	"github.com/mohuk/genie/dbase"
)

var (
	password = os.Getenv("DB_PWD")
	user     = os.Getenv("DB_USER")
	debug    = flag.Bool("debug", false, "enable debugging")
	port     = flag.Int("port", 1401, "the database port")
	host     = flag.String("host", "localhost", "the database host")
	database = flag.String("database", "WideWorldImporters", "the database name")
)

const (
	httpPort = 3000
)

func main() {

	if password == "" {
		password = "10Pearls!"
	}
	if user == "" {
		user = "sa"
	}

	store := dbase.NewStore(*host, *port, user, password)
	mapper := formly.NewFormlyMapper()
	gm := manager.NewGenieManager(store, mapper)

	APIRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	APIRouter.HandleFunc("/db", handlers.WithLogger(handlers.GetDatabases(gm))).Methods("GET")
	APIRouter.HandleFunc("/db/{dbname}/tables", handlers.GetTables(gm)).Methods("GET")
	APIRouter.HandleFunc("/db/{dbname}/tables/{tableId}", handlers.GetColumns(gm)).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), APIRouter))
}
