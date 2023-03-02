package mountain

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	inc, dec := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				inc[i] = max(inc[i], inc[j]+1)
			}
		}
	}

	for i := n - 2; i > -1; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				dec[i] = max(dec[i], dec[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if inc[i] > 0 && dec[i] > 0 {
			res = max(res, inc[i]+dec[i])
		}
	}

	return n - res - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
