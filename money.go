package gomoney

import (
	"strings"
)

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
// baseCurrency:    [string] 3 char value of string. Accepted values are (case insensitive)
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
	if !validCurrency(m.BaseCurrency) {
		return nil, ErrUnsupportedCurrency
	}
	er, err := newEurExchangeRate()
	if err != nil {
		return nil, err
	}
	m.exchangeRate = er
	return m, nil
}

// Converts the money into the currency inputted
// currency: [string] 3 char value of string. Accepted values are (case insensitive)
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
func (m *Money) Convert(currency string) (*Money, error) {
	currency = strings.ToUpper(currency)
	if !validCurrency(currency) {
		return nil, ErrUnsupportedCurrency
	}
	if !m.valid() {
		return nil, ErrInvalidMoneyObject
	}
	eurAmount := m.Amount / m.exchangeRate.Rates[m.BaseCurrency]
	return NewMoney(eurAmount*m.exchangeRate.Rates[currency], currency)
}

// Returns if the money struct is a valid one
func (m *Money) valid() bool {
	_, ok := supportedCurrencies[m.BaseCurrency]
	return m.exchangeRate != nil && ok
}
