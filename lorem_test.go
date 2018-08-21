package faker

import (
	"fmt"
	"strings"
	"testing"
)

func TestLorem(t *testing.T) {
	l := New().Lorem()
	Expect(t, "faker.Lorem", fmt.Sprintf("%T", l))
}

func TestWord(t *testing.T) {
	l := New().Lorem()
	Expect(t, true, len(l.Word()) > 0)
}

func TestWords(t *testing.T) {
	l := New().Lorem()
	Expect(t, 2, len(l.Words(2)))
}

func TestSentence(t *testing.T) {
	l := New().Lorem()
	split := strings.Split(l.Sentence(2), " ")
	Expect(t, 2, len(split))
}

func TestSentences(t *testing.T) {
	l := New().Lorem()
	sentences := l.Sentences(2)
	Expect(t, 2, len(sentences))
}

func TestParagraph(t *testing.T) {
	l := New().Lorem()
	split := strings.Split(l.Paragraph(2), ".")
	Expect(t, 3, len(split))
}

func TestParagraphs(t *testing.T) {
	l := New().Lorem()
	split := l.Paragraphs(2)
	Expect(t, 2, len(split))
}

func TestText(t *testing.T) {
	l := New().Lorem()
	text := l.Text(255)
	Expect(t, true, len(text) <= 255)
}

func TestBytes(t *testing.T) {
	l := New().Lorem()
	text := l.Bytes(255)
	Expect(t, true, len(text) <= 255)
}
