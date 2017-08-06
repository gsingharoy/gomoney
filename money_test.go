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

func TestMoneyConvert(t *testing.T) {
	t.Log("When invalid currency is passed")
	for _, curr := range []string{"xxx", "yyy", "345", "galleon"} {
		m := Money{Amount: 100, BaseCurrency: "EUR"}
		_, err := m.Convert(curr)
		assert.Equal(t, ErrUnsupportedCurrency, err, "Expected the currency to be unsupported")
	}

	t.Log("When the money struct is invalid")
	m1 := Money{BaseCurrency: "invalid curr"}
	_, err := m1.Convert("EUR")
	assert.Equal(t, ErrInvalidMoneyObject, err, "Expected invalid money object error")

	m2 := Money{BaseCurrency: "EUR"}
	_, err = m2.Convert("USD")
	assert.Equal(t, ErrInvalidMoneyObject, err, "Expected invalid money object error")

	t.Log("When the money struct is valid and a valid currency is passed")
	m3 := Money{BaseCurrency: "EUR", Amount: 10, exchangeRate: &eurExchangeRate{Rates: map[string]float64{"EUR": 1}}}
	amt, err := m3.Convert("EUR")
	assert.Nil(t, err)
	assert.Equal(t, float64(10), amt, "Expected amount to be the same")

	m4 := Money{BaseCurrency: "USD", Amount: 100, exchangeRate: &eurExchangeRate{Rates: map[string]float64{"EUR": 1, "USD": 1.1, "CHF": 0.9}}}
	amt, err = m4.Convert("CHF")
	assert.Nil(t, err)
	expectedAmount := float64((100 / 1.1) * 0.9)
	assert.Equal(t, expectedAmount, amt, "Expected amount incorrectly converted")
}
