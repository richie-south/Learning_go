package firstClass

// MyFilter comment
func MyFilter(array []string, fn func(string) bool) []string {
	result := []string{}
	for _, value := range array {
		if fn(value) {
			result = append(result, value)
		}
	}

	return result
}

// MyMap comment
func MyMap(array []string, fn func(string) string) []string {
	result := []string{}

	for i, value := range array {
		result = append(result[:i], fn(value))
	}
	return result
}
