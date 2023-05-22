package price

// totalPrice Compute and return the total price of hotel booking
// all amounts in input must be multiplied by 100. Currency is Dollar
// the amount returned must be divided by 100. (eg: 10132 => 101.32)
func totalPrice(nights, rate, cityTax uint) uint {
	return nights * (rate + cityTax)
}
