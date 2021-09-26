package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	words := strings.Fields(text)
	if len(words) < 1 {
		return nil
	}
	type wordCounts struct {
		word  string
		count int
	}
	wordsList := make([]wordCounts, 0)
	var result []string

	for _, wordOutside := range words {

		isPast := false
		for _, w := range wordsList {
			if w.word == wordOutside {
				isPast = true
				break
			}
		}
		if isPast {
			continue
		}

		currentWord := wordCounts{wordOutside, 0}
		for _, wordInside := range words {
			if wordInside == wordOutside {
				currentWord.count++
			}
		}
		wordsList = append(wordsList, currentWord)
	}

	sort.SliceStable(wordsList, func(i, j int) bool {
		if wordsList[i].count == wordsList[j].count {
			return wordsList[i].word < wordsList[j].word
		}
		return wordsList[i].count > wordsList[j].count
	})

	frequent := wordsList[:10]
	for i := 0; i < 10; i++ {
		result = append(result, frequent[i].word)
	}
	return result
}
