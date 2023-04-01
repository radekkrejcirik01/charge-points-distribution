package main

import (
	"log"

	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/rest"
)

func main() {
	database.Connect()

	if err := rest.Create().Listen(":8081"); err != nil {
		log.Fatal(err)
	}
}
