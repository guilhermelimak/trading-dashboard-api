package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var url = "https://mercadobitcoin.net/api"

func get(currency string, method string) string {
	res, err := http.Get(fmt.Sprintf("%s/%s/%s", url, currency, method))

	if err != nil {
		panic(err)
	}

	buf, _ := ioutil.ReadAll(res.Body)

	return string(buf)
}

func mercadoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	method := vars["method"]

	if method == "" {
		panic("You need a method (ticker, orderbook or trades)")
	}

	coin := vars["coin"]

	if coin == "" {
		panic("You need a coin (LTC, BTC, BCH)")
	}

	t := get(coin, method)

	fmt.Fprint(w, string(t))
}
