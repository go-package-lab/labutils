package labutils

import (
	"sort"
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

// IsNumeric returns true if the given character is a numeric, otherwise false.
func IsNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsHexDecimal returns true if the given character is a hexdecimal, otherwise false.
func IsHexDecimal(c byte) bool {
	return c >= '0' && c <= '9' || c >= 'a' && c <= 'f' || c >= 'A' && c <= 'F'
}

// IsAlphaNumeric returns true if the given character is a alphabet or a numeric, otherwise false.
func IsAlphaNumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

// A PrioritizedValue struct holds pair of an arbitrary value and a priority.
type PrioritizedValue struct {
	// Value is an arbitrary value that you want to prioritize.
	Value interface{}
	// Priority is a priority of the value.
	Priority int
}

// PrioritizedSlice is a slice of the PrioritizedValues
type PrioritizedSlice []PrioritizedValue

// Sort sorts the PrioritizedSlice in ascending order.
func (s PrioritizedSlice) Sort() {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Priority < s[j].Priority
	})
}

// Remove removes the given value from this slice.
func (s PrioritizedSlice) Remove(v interface{}) PrioritizedSlice {
	i := 0
	found := false
	for ; i < len(s); i++ {
		if s[i].Value == v {
			found = true
			break
		}
	}
	if !found {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
