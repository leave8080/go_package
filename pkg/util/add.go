package util

import "math"

// SumInt 求和整型
func SumInt(values ...int) int {
	var s = 0
	for _, val := range values {
		s += val
	}
	return s
}

// SumFloat64 求和浮点数float64
func SumFloat64(values ...float64) float64 {
	var s = 0.0
	for _, val := range values {
		s += val
	}
	return s
}

// FloatEquals 浮点数判等
func FloatEqual(a, b, precision float64) bool {
	if math.Abs(a-b) < precision {
		return true
	}
	return false
}
