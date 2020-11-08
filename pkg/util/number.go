package util

import (
	"fmt"
	"math"
	"strings"
)

// FormatNumber2String :
func FormatNumber2String(numberInt64 int64) string {
	number := float64(numberInt64)
	k := float64(1000)
	m := float64(1000000)
	b := float64(1000000000)
	t := float64(1000000000000)

	numStr := "0"
	if number < k {
		numStr = fmt.Sprintf("%.f", number)
	}
	if number >= k && number < m {
		remainder := math.Remainder(number, k)
		numStr = fmt.Sprintf("%.fK", math.Round(number/k))
		if remainder > 1 {
			numStr = fmt.Sprintf("%s.%.fK", strings.Split(numStr, "K")[0], math.Round(remainder*10/k))
		}
	}
	if number >= m && number < b {
		remainder := math.Remainder(number, m)
		numStr = fmt.Sprintf("%.fM", math.Round(number)/m)
		if remainder > 1 {
			numStr = fmt.Sprintf("%s.%.fM", strings.Split(numStr, "M")[0], math.Round(remainder*10/m))
		}
	}
	if number >= b && number < t {
		remainder := math.Remainder(number, b)
		numStr = fmt.Sprintf("%.fB", math.Round(number/b))
		if remainder > 1 {
			numStr = fmt.Sprintf("%s.%.fB", strings.Split(numStr, "B")[0], math.Round(remainder*10/b))
		}
	}
	if number >= t {
		remainder := math.Remainder(number, t)
		numStr = fmt.Sprintf("%.fT", math.Round(number)/t)
		if remainder > 1 {
			numStr = fmt.Sprintf("%s.%.fT", strings.Split(numStr, "T")[0], math.Round(remainder*10/t))
		}
	}

	return numStr
}
