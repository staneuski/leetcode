func checkSubarraySum(nums []int, k int) bool {
	remainderToIndex := make(map[int]int)
	remainderToIndex[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		remainder := sum % k
		value, exists := remainderToIndex[remainder]
		if exists && i-value >= 2 {
			return true
		} else if !exists {
			remainderToIndex[remainder] = i
		}
	}
	return false
}
