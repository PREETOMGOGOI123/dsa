package sqrt

func MySqrt(x int) int {
	low := 0
	high := x
	ans := 0

	for low <= high {
		mid := (low + high) / 2
		if mid*mid <= x {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
