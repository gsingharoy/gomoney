package gomoney

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Represents the exchange rate in terms of EUR
type eurExchangeRate struct {
	Date  time.Time          // Date on which enchange rate is taken
	Rates map[string]float64 // Map of exchange rates
}

// gets the exchange rate w.r.t. EUR from the european central bank SOAP endpoint
func newEurExchangeRate() (*eurExchangeRate, error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	r, err := httpClient.Get(urlEUExchangeRate)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return parseSOAPResponse(data)
}
