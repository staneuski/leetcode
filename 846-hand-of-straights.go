/**
 * 846-hand-of-straights.go
 *
 * https://leetcode.com/problems/hand-of-straights/
 * https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
 */

import "math"

func isNStraightHand(nums []int, groupSize int) bool {
	if len(nums)%groupSize != 0 {
		return false
	}

	for iGroup := 0; iGroup < len(nums)/groupSize; iGroup++ {
		i, num := findMin(nums)
		nums[i] = math.MaxInt32
		num++
		for j := 1; j < groupSize; j++ {
			i = find(nums, num)
			if i == len(nums) {
				return false
			}
			nums[i] = math.MaxInt32
			num++
		}
	}
	return true
}

func findMin(elems []int) (int, int) {
	ind, min := 0, math.MaxInt32
	for i, elem := range elems {
		if elem < min {
			ind, min = i, elem
		}
	}
	return ind, min
}

func find(elems []int, value int) int {
	for i, elem := range elems {
		if elem == value {
			return i
		}
	}
	return len(elems)
}
