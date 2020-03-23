package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	_, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Connected!")
	}
}
