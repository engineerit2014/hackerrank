package queues

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
	idx := int32(-1)
	accPetrol := int32(0)
	currentPetrol := int32(0)

	// evaluate each petrol and execute the tour
	for i := 0; i < len(petrolpumps); i++ {

		// if is a start point, store index
		if idx == -1 {
			idx = int32(i)
		}

		// if current petrol plus the petrol of the current pump is greater or equal to the distance
		if currentPetrol + petrolpumps[i][0] >= petrolpumps[i][1] {
			// then add the petrol to the current petrol accumulated
			currentPetrol += petrolpumps[i][0]
			// decrease the current distance and continue
			currentPetrol -= petrolpumps[i][1]
			continue
		}

		// if the condition is not valid, is because the petrol is not enough
		// the add all the petrol to an accumulated petrol
		accPetrol += currentPetrol
		accPetrol = petrolpumps[i][0] - petrolpumps[i][1]

		// reset current petrol and index
		currentPetrol = 0
		idx = -1

		// if is last elem, finish
		if i == len(petrolpumps) - 1 {
			break
		}
	}

	//fmt.Println(idx)
	return idx
}
