package utils

import "strings"

var (
	num2char = "0123456789abcdefghijklmnopqrstuvwxyz"
)

// NumToBHex ... 36
func NumToBHex(num int64) string {
	numStr := ""
	for num != 0 {
		yu := num % 36
		numStr = string(num2char[yu]) + numStr
		num = num / 36
	}
	return strings.ToUpper(numStr)
}
