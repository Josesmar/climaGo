package main

import (
	"clima-cep/internal/inbouind"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	inbouind.RegisterHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
