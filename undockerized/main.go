package main

import (
	"assignment2/handler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	//Starttime being initialized, will be used for statushandling
	startTime := time.Now()

	//Creating the port
	port := os.Getenv("PORT")
	//If the port cant be found, 8080 will be its port, and will then use that
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"

	}

	http.HandleFunc(handler.STATUS_PATH, func(w http.ResponseWriter, r *http.Request) {
		handler.MyStatusHandler(w, r, startTime)
	})

	//Starts and listens to the port
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
