package utils

import (
	"slices"
	"strings"
)

// func includes[value comparable, array []string | []int](value, array) bool {
func Includes(value string, array []string) bool {
	return slices.Contains(array, value)
}

// Check if the given string starts with any one of the strings in the given array of strings
func StartsWith(value string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(value, prefix) {
			return true
		}
	}
	return false
}

// Check if the given pointer is nill. If yes, return the placeholder, otherwise, return the value of the variable.
func GetValue(value *string, placeholder string) string {
	if value == nil {
		return placeholder
	}
	return *value
}
