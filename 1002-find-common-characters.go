import "math"

func commonChars(words []string) []string {
	totalCounts := make([]int, 26)
	for i := range totalCounts {
		totalCounts[i] = math.MaxInt32
	}

	for _, word := range words {
		wordCounts := countEach(word)
		for i, _ := range totalCounts {
			totalCounts[i] = min(totalCounts[i], wordCounts[i])
		}
	}

	return convertFromAscii(totalCounts)
}

func countEach(s string) []int {
	counts := make([]int, 26)
	for _, char := range s {
		counts[char-'a']++
	}
	return counts
}

func convertFromAscii(counts []int) []string {
	var res []string
	for i, count := range counts {
		for count > 0 {
			res = append(res, string(i+'a'))
			count--
		}
	}
	return res
}
