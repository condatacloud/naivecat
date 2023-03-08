package tools

import "math"

const MIN = 0.000001

// IsEqual
// 判断 float 是否相等
func IsEqual(x, y float64) bool {
	return math.Abs(x-y) < MIN
}
