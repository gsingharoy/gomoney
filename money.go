package gomoney

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
