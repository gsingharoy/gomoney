package gomoney

import "strings"

// Represents the basic Money unit.
// Contains 3 main fields, viz.,
// Amount, BaseCurrency and CurrConversion
type Money struct {
	// Amount which needs to be converted
	Amount float64
	// Base currency of the amount.
	BaseCurrency string
	// Exchange rate w.r.t. EUR
	exchangeRate *eurExchangeRate
}

// Returns a new Money instance
// amount:          [float64] Amount of the currency you want the money struct to be initialted with
// baseCurrency:    [string] 3 char value of string. Accepted values are
// []string{
//     "USD",
//     "JPY",
//     "BGN",
//     "CZK",
//     "DKK",
//     "GBP",
//     "HUF",
//     "PLN",
//     "RON",
//     "SEK",
//     "CHF",
//     "NOK",
//     "HRK",
//     "RUB",
//     "TRY",
//     "AUD",
//     "BRL",
//     "CAD",
//     "CNY",
//     "HKD",
//     "IDR",
//     "ILS",
//     "INR",
//     "KRW",
//     "MXN",
//     "MYR",
//     "NZD",
//     "PHP",
//     "SGD",
//     "THB",
//     "ZAR",
// }
func NewMoney(amount float64, baseCurrency string) (*Money, error) {
	m := &Money{Amount: amount, BaseCurrency: strings.ToUpper(baseCurrency)}
	_, ok := supportedCurrencies[m.BaseCurrency]
	if !ok {
		return nil, ErrUnsupportedCurrency
	}
	exchangeRate, err := newEurExchangeRate()
	if err != nil {
		return nil, err
	}
	m.exchangeRate = exchangeRate
	return m, nil
}
