package gameMath

// Sign returns -1 if value is less than zero.
//1 if value is more than zero.
//0 if value is zero.
func Sign(value float64) int {
	if value < 0 {
		return -1
	} else if value > 0 {
		return 1
	} else {
		return 0
	}
}
