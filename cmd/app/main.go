package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mdalasini/rssagg/internal/routes"
)

func check(err error, message string){
	if err != nil {
		log.Fatalln(message, err)
	}
}

func main(){
	err := godotenv.Load()
	check(err, "Error loading .env file: ")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT not found")
	}

	rounter := routes.NewRounter()

	srv := &http.Server{
		Handler: rounter,
		Addr: ":" + portString,
	}

	log.Printf("Server stating on port %v", portString)

	err = srv.ListenAndServe()
	check(err, "Error listening and servering server: ")
}