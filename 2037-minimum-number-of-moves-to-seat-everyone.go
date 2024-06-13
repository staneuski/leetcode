/**
 * 2037-minimum-number-of-moves-to-seat-everyone.go
 *
 * https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone
 */

import "slices"

func minMovesToSeat(seats []int, students []int) int {
	slices.Sort(seats)
	slices.Sort(students)

	// Greedily match each student to a seat.
	// The smallest positioned student will go to the smallest positioned chair,
	// and then the next smallest positioned student will go to 
	// the next smallest positioned chair
	moveCount := 0
	for i, seat := range seats {
		moveCount += abs(seat - students[i])
	}
	return moveCount
}

func abs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}
