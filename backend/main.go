package main

import (
	"backend/meals"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

func main() {
	port := 8081

	db, err := sql.Open("sqlite", "./meals.db")
	if err != nil {
		log.Fatalf("error while opening database conneciton. Error: %s", err.Error())
	}

	createTable(db)

	ms := meals.NewSqliteMealsService(db)
	mh := meals.NewMealHandler(ms)

	router := mux.NewRouter().UseEncodedPath()

	mealsRouter := router.PathPrefix("/meals").Subrouter()
	mealsRouter.Methods("GET").Path("").HandlerFunc(mh.ListMeals)
	mealsRouter.Methods("POST").Path("").HandlerFunc(mh.CreateMeal)
	mealsRouter.Methods("DELETE").Path("/{id}").HandlerFunc(mh.DeleteMeal)
	mealsRouter.Methods("PUT").Path("").HandlerFunc(mh.EditMeal)

	suggestionsRouter := router.PathPrefix("/suggestions").Subrouter()
	suggestionsRouter.Path("").Methods("GET").HandlerFunc(mh.SuggestMeals)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	server := &http.Server{
		Handler:     c.Handler(router),
		Addr:        fmt.Sprintf(":%d", port),
		ReadTimeout: 10 * time.Second,
	}

	log.Infof("Starting server with port: %d\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting the server: %e", err.Error())
	}
}

func createTable(db *sql.DB) {
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS "meals" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" VARCHAR(64) NOT NULL,
		"category" VARCHAR(64) NOT NULL
	)`)

	if err != nil {
		log.Fatalf("unable to prepare create meals table statement. Error %s", err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("unable to execute create meals table statement. Error %s", err.Error())
	}
}
