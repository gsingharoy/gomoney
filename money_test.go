package gomoney

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMoney(t *testing.T) {
	t.Log("When the currency is an unsupported currency")
	for _, curr := range []string{"xxx", "yyy", "345", "galleon"} {
		_, err := NewMoney(100, curr)
		assert.Equal(t, ErrUnsupportedCurrency, err, "Expected the currency to be unsupported")
	}

	t.Log("When the currency is valid")
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
		expectedMoney := &Money{Amount: 42, BaseCurrency: strings.ToUpper(curr)}
		assert.Equal(t, expectedMoney, m, "the money struct does not match")
	}
}
