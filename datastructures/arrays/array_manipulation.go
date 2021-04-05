package arrays

// Complete the arrayManipulation function below.
func ArrayManipulation(n int32, queries [3][3]int32) int64 {
	nIncluded := n + 1
	result := make([]int64, nIncluded, nIncluded)

	for i := 0; i < len(queries); i++ {
		a := int64(queries[i][0])
		b := int64(queries[i][1])
		k := int64(queries[i][2])

		result[a-1] += k
		result[b] -= k
	}

	var max int64 = 0
	var sum int64 = 0

	// prefix sum
	for i := int32(0); i < nIncluded; i++ {
		sum += result[i]
		if sum > max {
			max = sum
		}
	}

	return max
}
