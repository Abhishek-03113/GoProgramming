package Conversions

func Length(meters float64) float64 {
	return meters * 3.281
}

func Temperature(celcius float64) float64 {
	return (9 * celcius / 5) + 32

}

func Weight(kilogram float64) float64 {
	return kilogram * 2.205
}
