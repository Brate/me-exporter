package helpers

import "strconv"

func StringToFloat(fStr string) float64 {
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
