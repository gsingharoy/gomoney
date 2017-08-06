package gomoney

import "strings"

// checks if the currency passed is a supported valid currency
func validCurrency(currency string) bool {
	_, ok := supportedCurrencies[strings.ToUpper(currency)]
	return ok
}
