/**
 * 648-replace-words.go
 *
 * https://leetcode.com/problems/replace-words/
 */

import "strings"

func replaceWords(prefixes []string, sentence string) string {
	var res strings.Builder
	res.Grow(len(sentence))
	for i, word := range strings.Split(sentence, " ") {
		if i != 0 {
			res.WriteString(" ")
		}
		res.WriteString(shortestPrefix(word, prefixes))
	}
	return res.String()
}

func shortestPrefix(s string, prefixes []string) string {
	shortest := s
	for _, prefix := range prefixes {
		if len(prefix) < len(shortest) && strings.HasPrefix(s, prefix) {
			shortest = prefix
		}
	}
	return shortest
}
