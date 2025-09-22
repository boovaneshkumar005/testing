package main

import (
	"gotesting/config"
	"gotesting/handler"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUserHandler(w, r)
		case http.MethodPost:
			handler.AddUserHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
