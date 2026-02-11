package operator

func Div(x, y float64) float64 {
	if y <= 0 {
		panic("divisor mustn't be zero or negative")
	}
	return x / y
}