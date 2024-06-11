/**
 * 1122-relative-sort-array.go
 *
 * https://leetcode.com/problems/relative-sort-array/
 */

import "slices"

func relativeSortArray(elems []int, order []int) []int {
	front := 0
	for _, num := range order {
		index := slices.IndexFunc(elems[front:], func(elem int) bool { return elem == num })
		for index != -1 {
			toFront(elems[front:], index)
			front++
			index = slices.IndexFunc(elems[front:], func(elem int) bool { return elem == num })
		}
	}
	slices.Sort(elems[front:])
	return elems
}

func toFront(elems []int, index int) {
	elems[0], elems[index] = elems[index], elems[0]
}
