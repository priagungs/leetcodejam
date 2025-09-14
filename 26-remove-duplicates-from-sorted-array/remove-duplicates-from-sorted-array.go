
func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	a, b, c := 0, 1, len(nums)
	for b < c {
		if nums[a] != nums[b] {
			nums = append(nums[:a+1], nums[b:]...)
			c = len(nums)
			a++
			b = a
		}
		b++
	}
    if b-a > 1 {
        nums = nums[:a+1]
    }
	return len(nums)
}
