package faker

import (
	_ "embed"
	"strings"
)

//go:embed data/english_words.txt
var englishWordsData string

var englishWords = parseEnglishWords(englishWordsData)

func parseEnglishWords(data string) []string {
	if data == "" {
		return nil
	}
	return strings.Split(strings.TrimSpace(data), "\n")
}
