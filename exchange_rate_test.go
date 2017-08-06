package gomoney

import (
	"io/ioutil"
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	"github.com/stretchr/testify/assert"
)

func TestNewEurExchangeRate(t *testing.T) {
	t.Log("When the endpoint is working")
	data, err := ioutil.ReadFile("./fixtures/exchange_rates/1.xml")
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("GET", "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml",
		httpmock.NewStringResponder(200, string(data)))

	m, err := newEurExchangeRate()
	assert.Nil(t, err, "Expected no error to be returned")
	expectedDate, _ := time.Parse("2006-01-02", "2017-08-04")
	assert.Equal(t, expectedDate, m.Date, "Exchange rate date does not match")
	expectedExchangeRate := map[string]float64{
		"EUR": 1,
		"USD": 1.1868,
		"JPY": 130.67,
		"BGN": 1.9558,
		"CZK": 26.068,
		"DKK": 7.4388,
		"GBP": 0.90280,
		"HUF": 304.36,
		"PLN": 4.2410,
		"RON": 4.5635,
		"SEK": 9.6053,
		"CHF": 1.1494,
		"NOK": 9.3618,
		"HRK": 7.4070,
		"RUB": 71.4691,
		"TRY": 4.1903,
		"AUD": 1.4888,
		"BRL": 3.6911,
		"CAD": 1.4920,
		"CNY": 7.9757,
		"HKD": 9.2782,
		"IDR": 15791.56,
		"ILS": 4.2939,
		"INR": 75.5035,
		"KRW": 1334.12,
		"MXN": 21.1344,
		"MYR": 5.0777,
		"NZD": 1.5928,
		"PHP": 59.588,
		"SGD": 1.6097,
		"THB": 39.455,
		"ZAR": 15.8267,
	}
	assert.Equal(t, expectedExchangeRate, m.Rates, "Exchange rates does not match with the expected type")
}
