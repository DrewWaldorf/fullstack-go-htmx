package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/fullstackgo/handlers"
	"github.com/sikozonpc/fullstackgo/store"
)

func main() {

	db, err := store.NewMySQLStorage()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(db)

	initStorage(db)

	router := mux.NewRouter()

	handler := handlers.New(store)

	router.HandleFunc("/", handler.HandleHome).Methods("GET")
	router.HandleFunc("/cars", handler.HandleListCars).Methods("GET")
	router.HandleFunc("/cars", handler.HandleAddCar).Methods("POST")
	router.HandleFunc("/cars/{id}", handler.HandleDeleteCar).Methods("DELETE")
	router.HandleFunc("/cars/search", handler.HandleSearchCar).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", "localhost:8080")
	http.ListenAndServe(":8080", router)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database!")
}
