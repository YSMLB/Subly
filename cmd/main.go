package main

import (
	"fmt"
	"log"
	"net/http"

	"Subly/internal/config"
	"Subly/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ананьев Чмо")
	})
	log.Println("Server is started on: 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
