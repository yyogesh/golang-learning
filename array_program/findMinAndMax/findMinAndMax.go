package findminandmax

func FindMinAndMax(arr []int) (int, int) {
	max, min := arr[0], arr[0]

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
