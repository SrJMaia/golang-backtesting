package conversion

import "math"

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

func Round(f float64, p float64) float64 {
	/*
		value := 3252.195
		fmt.Println(Round(value, 3)) // 3252.194
		fmt.Println(Round(value, 2)) // 3252.19
		fmt.Println(Round(value, 1)) // 3252.2
		fmt.Println(Round(value, 0)) // 3252
		fmt.Println(Round(value, -1)) // 3250
		fmt.Println(Round(value, -2)) // 3300
		fmt.Println(Round(value, -3)) // 3000
	*/
	var x float64 = math.Pow(10, p)
	var y float64 = 0.5 / x
	return math.Floor((f+y)*x) / x
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
