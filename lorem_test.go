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
	got := l.Word()
	Expect(t, true, len(got) > 0)
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

func TestTextUsesRandomWords(t *testing.T) {
	f1 := NewWithSeedInt64(42)
	f2 := NewWithSeedInt64(99)
	text1 := f1.Lorem().Text(50)
	text2 := f2.Lorem().Text(50)
	Expect(t, true, len(text1) > 0)
	Expect(t, true, len(text1) <= 50)
	Expect(t, true, text1 != text2)
	Expect(t, true, strings.Contains(text1, " "))
}

func TestBytes(t *testing.T) {
	l := New().Lorem()
	text := l.Bytes(255)
	Expect(t, true, len(text) <= 255)
}
