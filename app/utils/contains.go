package utils

// func includes[value comparable, array []string | []int](value, array) bool {
func Includes(value string, array []string) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}