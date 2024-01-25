package utility

// IndexesOf return every index where `target` exists in `array`
func IndexesOf[K comparable](target K, array []K) []int {
	var matchedIndexes []int
	for i, v := range array {
		if v == target {
			matchedIndexes = append(matchedIndexes, i)
		}
	}
	return matchedIndexes
}

// IndexOf returns the FIRST occurrence of `target` in `array`
//
// NOTE: Use `IndexesOf` to return ALL occurrences
func IndexOf[K comparable](target K, array []K) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}
