package arrays

import "fmt"

/*
 * Complete the 'dynamicArray' function below.
 * 2 2
 * 1 x y
 * 2 x y
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func DynamicArray(n int32, queries [][]int32) []int32 {
	// initialize inputs and output
	var result []int32
	arr := make([][]int32, n)
	numQueries := len(queries)
	lastAnswer := int32(0)

	// iterate each query
	for i := 0; i < numQueries; i++ {
		x := queries[i][1]
		y := queries[i][2]
		idx := (x ^ lastAnswer) % n

		// query type 1
		if queries[i][0] == 1 {
			arr[idx] = append(arr[idx], y)
		}

		// query type 2
		if queries[i][0] == 2 {
			lastAnswer = arr[idx][y%int32(len(arr[idx]))]
			result = append(result, lastAnswer)
			fmt.Println(lastAnswer)
		}
	}

	return result
}
