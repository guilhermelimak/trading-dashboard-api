package mercadobtc

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

// Handler is used when receiving a request to mercadobitcoin exchange
func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	method := vars["method"]
	if method == "" {
		fmt.Fprint(w, "You need a method (orderbook, trades, ticker)")
	}

	coin := vars["coin"]
	if coin == "" {
		fmt.Fprint(w, "You need a coin (LTC, BTC, BCH)")
	}

	t := get(coin, method)

	fmt.Fprint(w, string(t))
}
