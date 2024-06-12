/**
 * 75-sort-color.go
 *
 * https://leetcode.com/problems/sort-colors/
 */

import "slices"

func sortColors(elems []int) {
	// Similar to problem no. 1122. Relative Sort Array
	front := 0
	for color := 0; color < 2; color++ {
		index := slices.IndexFunc(elems[front:], func(elem int) bool { return elem == color })
		for index != -1 {
			toFront(elems[front:], index)
			front++
			index = slices.IndexFunc(elems[front:], func(elem int) bool { return elem == color })
		}
	}
}

func toFront(elems []int, index int) {
	elems[0], elems[index] = elems[index], elems[0]
}
