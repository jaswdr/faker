package faker

import (
	"testing"
)

func TestGetWord(t *testing.T) {
	tables := []struct {
		context     string
		mockFunc    func()
		expectWords []string
	}{
		{
			context: "test",
			mockFunc: func() {
				consultService = func(context string, max int, w *[]WordInfo) {
					*w = []WordInfo{
						{Word: "test1", Score: 1, Tags: []string{"tag1"}},
						{Word: "test2", Score: 2, Tags: []string{"tag2"}},
						{Word: "test3", Score: 3, Tags: []string{"tag3"}},
					}
				}
			},
			expectWords: []string{"test1", "test2", "test3"},
		},
	}

	originalResponse := consultService
	for _, table := range tables {
		table.mockFunc()
		resultWords := getWords(table.context)
		if len(resultWords) != len(table.expectWords) {
			t.Errorf("getWords(%s) was incorrect, got: %d, want: %d.", table.context, len(resultWords), len(table.expectWords))
		}
		resultWord := getWord(table.context)
		valid := false
		for _, w := range table.expectWords {
			if w == resultWord {
				valid = true
				break
			}
		}
		if !valid {
			t.Errorf("getWord(%s) was incorrect, got: %s.", table.context, resultWord)
		}
		consultService = originalResponse
	}
}
