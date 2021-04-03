package arrays

// Complete the hourglassSum function below.
func HourglassSum(arr [][]int32) int32 {
	var r int32
	var acc int32

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {

			acc = 0
			//build hourglass
			for y := i; y < i+3; y++ {
				for x := j; x < j+3; x++ {

					if y == i+1 && (x == j+0 || x == j+2) {
						acc += 0
					} else {
						acc += arr[y][x]
					}

				}
			}

			if i == 0 && j == 0 {
				r = acc
			}

			if acc > r {
				r = acc
			}
		}
	}

	return r
}
