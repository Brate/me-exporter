package helpers

import "strconv"

func StringToFloat(string fStr) float {
	f, err := strconv.ParseFloat(fStr, 64)
	if err != nil {
		return 0.0
	}
	return f
}
