package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Docker!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Docker world")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
