package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func main() {
	// this Pings the database trying to connect, panics on error
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()

	// Try curl localhost:8000/people | jq
	router.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		people := []Person{}
		db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
		json.NewEncoder(w).Encode(people)
	})

	// Try curl localhost:8000/people/jane.citzen@example.com | jq
	router.HandleFunc("/people/{email}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email := vars["email"]
		person := Person{}
		err = db.Get(&person, "SELECT * FROM person WHERE email=$1", email)
		json.NewEncoder(w).Encode(person)
	})

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
