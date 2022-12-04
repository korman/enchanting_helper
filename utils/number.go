package utils

import "strconv"

func N2S(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 2, 64)
}
