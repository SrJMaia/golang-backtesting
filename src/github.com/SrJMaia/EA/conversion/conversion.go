package conversion

func BoolToFloat(b bool) float64 {
	if b {
		return 1.
	}
	return 0.
}

func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func FloatToBool(f float64) bool {
	if f == 1. {
		return true
	}
	return false
}

func IntToBool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}
