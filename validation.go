package labutils

import (
	"unicode"
)

// IsInteger 判断字符串是否全是整数
func IsInteger(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) == false {
			return false
		}
	}
	return true
}
