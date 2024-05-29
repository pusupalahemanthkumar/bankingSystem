package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// Util function to check supported currencies
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false

}
