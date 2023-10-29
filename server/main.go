package main

import (
	"bin_components"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	component := bin_components.Index()

	router := mux.NewRouter()
	router.Handle("/", templ.Handler(component))

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on port :8080")
	log.Fatal(server.ListenAndServe())
}
