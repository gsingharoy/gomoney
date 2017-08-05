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
	// Converted values of the amount in various currencies
	currConversion *currencyConversion
}

func NewMoney(amount float64, baseCurrency string) (*Money, error) {
	m := &Money{Amount: amount, BaseCurrency: strings.ToUpper(baseCurrency)}
	_, ok := supportedCurrencies[m.BaseCurrency]
	if !ok {
		return nil, ErrUnsupportedCurrency
	}
	return m, nil
}
