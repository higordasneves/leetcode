package sorted_squares

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	for i := len(result) - 1; i >= 0; i-- {
		leftItem := nums[0]
		index := len(nums) - 1
		rightItem := nums[index]

		if abs(leftItem) > abs(rightItem) {
			result[i] = leftItem * leftItem
			nums = nums[1:]
		} else {
			result[i] = rightItem * rightItem
			nums = nums[:index]
		}
	}

	return result
}

func sortedSquaresV2(nums []int) []int {
	result := make([]int, len(nums))

	for n := len(nums); n > 0; n = len(nums) {
		leftItem := nums[0]
		index := n - 1
		rightItem := nums[index]

		if abs(leftItem) > abs(rightItem) {
			nums = nums[1:]
			result[index] = leftItem * leftItem
		} else {
			nums = nums[:index]
			result[index] = rightItem * rightItem
		}
	}

	return result
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
