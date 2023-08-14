package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scw-be/pkg/api"
	"scw-be/pkg/db"
)

func main() {
	log.Print("server running")

	//start the db
	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	//get the router of the API by passing the db
	router := api.StartAPI(pgdb)

	//get the port from the env variables
	port := os.Getenv("PORT")

	//pass the router and start listening with the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
