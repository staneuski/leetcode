/**
 * 409-longest-palindrome.go
 * 
 * https://leetcode.com/problems/longest-palindrome
 */

func longestPalindrome(s string) int {
	length := 0
	occurences := map[rune]int{}
	for _, c := range s {
		occurences[c]++
		if occurences[c]%2 == 0 {
			length += 2
		}
	}
	if length == len(s) {
		return len(s)
	}
	return length + 1
}
