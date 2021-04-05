package queues

import "fmt"

/*
 * truckTour: There are  petrol pumps at a circle. Petrol pumps are numbered 1 to N-1 (both inclusive)
 *
 * Considerations:
 * 1. You have a Tank without petrol to do the tour.
 * 2. Start the tour at any of the Petrol Pumps.
 * 3. Consider that the truck will stop at each of the petrol pumps.
 * 4. The truck will move one kilometer for each litre of the petrol.
 *
 * Input: Each petrol pump is a row of the matrix:
 * Row -> (the amount of petrol that petrol pump will give,  the distance between that petrol pump and the next petrol pump)
 *
 * Output: Calculate the first point from where the truck will be able to complete the circle.
 * An integer which will be the smallest index of the petrol pump from which we can start the tour.
 */
func TruckTour(petrolpumps [][]int32) int32 {
	// evaluate if each petrol is a valid start point and execute the tour
	for i := 0; i < len(petrolpumps); i++ {

		// if amount of petrol is greater or equal to the distance, then is a possible start point
		if petrolpumps[i][0] >= petrolpumps[i][1] {

			// put the selected start point as head of the queue
			// validate the tour with the start point selected
			if isValidTour(append(petrolpumps[i:], petrolpumps[0:i]...)) {
				fmt.Println(i)
				return int32(i)
			}
		}
	}

	return 0
}

func isValidTour(queue [][]int32) bool {
	// mark start point
	startPoint := make([]int32, 2)
	startPoint[0] = queue[0][0]
	startPoint[1] = queue[0][1]

	for i := 1; i < len(queue); i++ {
		// calculate the current petrol (petrol - distance)
		currentP := startPoint[0] - startPoint[1]
		if currentP < 0 {
			return false
		}

		// calculate next petrol: currentP + petrolNextPoint
		currentP += queue[i][0]

		// if current petrol is greater or equal to the next distance, then select a new start point
		if currentP < queue[i][1] {
			return false
		}
		startPoint[0] = currentP
		startPoint[1] = queue[i][1]
	}

	return true
}
