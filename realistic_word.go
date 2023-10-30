package faker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type WordInfo struct {
	Word  string   `json:"word"`
	Score int      `json:"score"`
	Tags  []string `json:"tags"`
}

const apiEndpoint = "https://api.datamuse.com/words?ml=%s&max=%d"

var consultService = func(context string, max int, w *[]WordInfo) {
	replaceContext := strings.Replace(context, " ", "+", -1)
	apiURL := fmt.Sprintf(apiEndpoint, replaceContext, max)

	resData, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error in HTTP request:", err)
		return
	}
	defer resData.Body.Close()

	bodyRes, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		fmt.Println("Error reading response data:", err)
		return
	}

	err = json.Unmarshal([]byte(bodyRes), &w)
	if err != nil {
		fmt.Println("Error deserializing JSON:", err)
		return
	}
}

func getWords(context string) []string {
	var words []WordInfo
	consultService(context, 20, &words)

	result := []string{}
	for _, word := range words {
		result = append(result, word.Word)
	}
	return result
}

func getRandomIndex(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max + 1)
}

func getWord(context string) string {
	var words []WordInfo
	consultService(context, 10, &words)
	randIndex := getRandomIndex(len(words) - 1)
	return words[randIndex].Word
}
