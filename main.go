package main

import (
	"fmt"
	"net/http"
	"trading-dashboard-api/mercadobtc"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Port the server will listen on.
const Port = 8888

// Address the server will listen on.
const Address = "0.0.0.0"

// ServerAddress formatted on address:port format.
var ServerAddress = fmt.Sprintf("%s:%d", Address, Port)

// ClientAddress to be allowed on cors.
const ClientAddress = "http://localhost:8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mercadobtc/{coin}/{method}", mercadobtc.Handler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{ClientAddress},
		AllowCredentials: true,
	})

	http.Handle("/", c.Handler(router))

	fmt.Printf("Listening on: %s...", ServerAddress)

	err := http.ListenAndServe(ServerAddress, nil)

	if err != nil {
		panic(err)
	}
}
