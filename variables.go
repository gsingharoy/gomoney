package gomoney

import "errors"

// Standard currency values supported by the library
var supportedCurrencies = map[string]bool{
	"EUR": true,
	"USD": true,
	"JPY": true,
	"BGN": true,
	"CZK": true,
	"DKK": true,
	"GBP": true,
	"HUF": true,
	"PLN": true,
	"RON": true,
	"SEK": true,
	"CHF": true,
	"NOK": true,
	"HRK": true,
	"RUB": true,
	"TRY": true,
	"AUD": true,
	"BRL": true,
	"CAD": true,
	"CNY": true,
	"HKD": true,
	"IDR": true,
	"ILS": true,
	"INR": true,
	"KRW": true,
	"MXN": true,
	"MYR": true,
	"NZD": true,
	"PHP": true,
	"SGD": true,
	"THB": true,
	"ZAR": true,
}

// Define all errors here

var ErrUnsupportedCurrency = errors.New("unsupported currency")
var ErrInvalidMoneyObject = errors.New("missing exchange rates")
