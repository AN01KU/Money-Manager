package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	var r *chi.Mux = chi.NewRouter()
	fmt.Println("Starting GO API Service ...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
