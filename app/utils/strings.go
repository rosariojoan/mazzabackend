package utils

import "strings"

// func includes[value comparable, array []string | []int](value, array) bool {
func Includes(value string, array []string) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
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
