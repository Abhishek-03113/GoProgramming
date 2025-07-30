package ExpressionEvaluator

func Add(a, b float64) float64 {
	return a + b
}

func Minus(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	return a / b
}

func Modulo(a, b float64) int {
	return int(a) % int(b)
}

func Equals(a, b float64) bool {
	return a == b
}

func NotEquals(a, b float64) bool {
	return !Equals(a, b)

}

func GreaterThan(a, b float64) bool {
	return a > b
}

func GreaterThanEquals(a, b float64) bool {
	return GreaterThan(a, b) || a == b

}

func LessThan(a, b float64) bool {
	return a < b
}

func LessThanEquals(a, b float64) bool {
	return LessThan(a, b) || a == b

}
