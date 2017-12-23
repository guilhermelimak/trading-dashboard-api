package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mercadobtc/{coin}/{method}", mercadoHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	http.Handle("/", c.Handler(router))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
