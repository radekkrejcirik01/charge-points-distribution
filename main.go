package main

import (
	"fmt"
	"log"

	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/rest"
)

func main() {
	fmt.Println("App is running!")
	database.Connect()

	if err := rest.Create().Listen(":8081"); err != nil {
		log.Fatal(err)
	}
}
