package program

func FindMaxandMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	// {0, 1, 2, 3, 4, 5}

	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}
