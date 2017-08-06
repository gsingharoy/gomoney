package gomoney

import (
	"io/ioutil"
	"strings"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	"github.com/stretchr/testify/assert"
)

func TestNewMoney(t *testing.T) {
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
	httpmock.RegisterResponder("GET", "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml",
		httpmock.NewStringResponder(200, string(data)))
	validCurrs := []string{
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
	for _, curr := range validCurrs {
		m, err := NewMoney(42, strings.ToLower(curr))
		assert.Nil(t, err, "Expected no error")
		expectedMoney := &Money{Amount: 42, BaseCurrency: strings.ToUpper(curr), exchangeRate: expectedEurExchangeRate1()}
		assert.Equal(t, expectedMoney, m, "the money struct does not match")
	}
}
