package arrays

// Complete the matchingStrings function below.
func MatchingStrings(strings []string, queries []string) []int32 {
	result := make([]int32, len(queries))

	for i, q := range queries {
		result[i] = 0
		for _, str := range strings {
			if str == q {
				result[i]++
			}
		}
	}

	return result
}
