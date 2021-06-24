package conversion

import (
	"strconv"
)

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

func Round(f float64, p int) float64 {
	y := strconv.FormatFloat(f, 'f', p, 64)
	r, _ := strconv.ParseFloat(y, 64)
	return r
}

func RemoveExcessZeros(s []float64) []float64 {
	var breakPoint int
	for i := range s {
		if s[i] == 0 {
			breakPoint = i
			break
		}
	}
	return s[:breakPoint]
}
