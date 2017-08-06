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
	// check if result is cached and try to fetch from memory first
	cResult, ok := cache.Find(cacheKeyexchangeRate)
	if ok {
		return cResult.(*eurExchangeRate), nil
	}

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
	result, err := parseSOAPResponse(data)
	if err != nil {
		return nil, err
	}
	go cache.Set(cacheKeyexchangeRate, result, 7200) // Result cached for 2 hours
	return result, nil
}
