package main

import (
	"fmt"
	"learn-go/db"
	"learn-go/routes"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := http.NewServeMux()

	db := db.InitDB()
	defer db.Close()

	routes.UserRoutes(router, db)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello developer, you've requested is coming from: %s\n", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static/"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8080...")
	server.ListenAndServe()
}
