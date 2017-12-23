package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mercado/{coin}/{method}", mercadoHandler)

	http.Handle("/", r)
	err := http.ListenAndServe(":6666", nil)

	if err != nil {
		panic(err)
	}
}
