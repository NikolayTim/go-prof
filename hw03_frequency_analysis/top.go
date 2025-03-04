package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var pattern = regexp.MustCompile(`\s`)

func Top10(str string) []string {
	if str == "" {
		return nil
	}

	type sortedWord struct {
		word  string
		count int
	}

	words := strings.Split(
		string(pattern.ReplaceAll([]byte(str), []byte(" "))),
		" ")
	wordsMap := make(map[string]int, len(words))

	for _, v := range words {
		if v == "" {
			continue
		}

		wordsMap[v]++
	}

	var sortedWords []sortedWord
	for word, count := range wordsMap {
		sortedWords = append(sortedWords, sortedWord{word, count})
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		if sortedWords[i].count != sortedWords[j].count {
			return sortedWords[i].count > sortedWords[j].count
		}

		return sortedWords[i].word < sortedWords[j].word
	})

	var result []string
	for i, v := range sortedWords {
		if i > 9 {
			break
		}

		result = append(result, v.word)
	}

	return result
}
