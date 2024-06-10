/**
 * 1051-height-checker.go
 *
 * https://leetcode.com/problems/height-checker
 */

import "sort"

func heightChecker(heights []int) int {
    expected := make([]int, len(heights))
    copy(expected, heights)
    sort.Ints(expected)

    count := 0
    for i, height := range heights {
        if height != expected[i] {
            count++
        }
    }
    return count
}
