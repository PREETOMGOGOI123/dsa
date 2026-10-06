package plusOne

func PlusOne(digits []int) []int {

	lod := len(digits)
	for i := lod - 1; i >= 0; i-- {
		if i-1 < 0 {
			if digits[i] == 9 {
				digits[i] = 0
				digits = append([]int{1}, digits...)
				return digits
			}
			digits[i]++
			return digits
		}
		// when the next posittion exists
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	return digits

}
