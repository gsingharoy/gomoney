package gomoney

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestNewMoney(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	t.Log("When the currency is an unsupported currency")
	for _, curr := range []string{"xxx", "yyy", "345", "galleon"} {
		_, err := NewMoney(100, curr)
		assert.Equal(t, ErrUnsupportedCurrency, err, "Expected the currency to be unsupported")
	}

	t.Log("When the currency is valid")
	data, err := ioutil.ReadFile("./fixtures/exchange_rates/1.xml")
	if err != nil {
		panic(err)
	}
	httpmock.RegisterResponder("GET", urlEUExchangeRate,
		httpmock.NewStringResponder(200, string(data)))
	validCurrs := []string{
		"EUR",
		"USD",
		"JPY",
		"BGN",
		"CZK",
		"DKK",
		"GBP",
		"HUF",
		"PLN",
		"RON",
		"SEK",
		"CHF",
		"NOK",
		"HRK",
		"RUB",
		"TRY",
		"AUD",
		"BRL",
		"CAD",
		"CNY",
		"HKD",
		"IDR",
		"ILS",
		"INR",
		"KRW",
		"MXN",
		"MYR",
		"NZD",
		"PHP",
		"SGD",
		"THB",
		"ZAR",
	}
	expectedExchangeRate := expectedEurExchangeRate1()
	for _, curr := range validCurrs {
		m, err := NewMoney(42, strings.ToLower(curr))
		assert.Nil(t, err, "Expected no error")
		assert.Equal(t, &Money{Amount: float64(42), BaseCurrency: strings.ToUpper(curr), exchangeRate: expectedExchangeRate}, m, "the money struct does not match")
	}
}
