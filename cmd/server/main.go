package main

import (
	"fmt"
	"net/http"

	"github.com/AN01KU/money-manager/internal/handlers"
	"github.com/AN01KU/money-manager/internal/tools"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Load .env file
	godotenv.Load()

	var err error
	var r *chi.Mux = chi.NewRouter()
	fmt.Println("Starting GO API Service ...")

	// setup DB
	database, err := tools.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return
	}

	handlers := handlers.NewHandler(database)
	handlers.RegisterRoutes(r)

	err = http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
