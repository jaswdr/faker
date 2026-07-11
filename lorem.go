package faker

import (
	"strings"
)

// Lorem is a faker struct for Lorem
type Lorem struct {
	Faker *Faker
}

// Word returns a fake word for Lorem
func (l Lorem) Word() string {
	index := l.Faker.IntBetween(0, len(englishWords)-1)
	return englishWords[index]
}

// Words returns fake words for Lorem
func (l Lorem) Words(nbWords int) []string {
	words := make([]string, 0, nbWords)
	for i := 0; i < nbWords; i++ {
		words = append(words, l.Word())
	}

	return words
}

// Sentence returns a fake sentence for Lorem
func (l Lorem) Sentence(nbWords int) string {
	return strings.Join(l.Words(nbWords), " ") + "."
}

// Sentences returns fake sentences for Lorem
func (l Lorem) Sentences(nbSentences int) []string {
	sentences := make([]string, 0, nbSentences)
	for i := 0; i < nbSentences; i++ {
		sentences = append(sentences, l.Sentence(l.Faker.RandomNumber(2)))
	}

	return sentences
}

// Paragraph returns a fake paragraph for Lorem
func (l Lorem) Paragraph(nbSentences int) string {
	return strings.Join(l.Sentences(nbSentences), " ")
}

// Paragraphs returns fake paragraphs for Lorem
func (l Lorem) Paragraphs(nbParagraph int) []string {
	out := make([]string, 0, nbParagraph)
	for i := 0; i < nbParagraph; i++ {
		out = append(out, l.Paragraph(l.Faker.RandomNumber(2)))
	}

	return out
}

// Text returns a fake text for Lorem using randomly selected words up to maxNbChars.
func (l Lorem) Text(maxNbChars int) string {
	if maxNbChars <= 0 {
		return ""
	}

	var builder strings.Builder
	for builder.Len() < maxNbChars {
		word := l.Word()
		if builder.Len() > 0 {
			if builder.Len()+1+len(word) > maxNbChars {
				break
			}
			builder.WriteByte(' ')
		} else if len(word) > maxNbChars {
			return word[:maxNbChars]
		}
		builder.WriteString(word)
	}

	return builder.String()
}

// Bytes returns fake bytes for Lorem
func (l Lorem) Bytes(maxNbChars int) (out []byte) {
	return []byte(l.Text(maxNbChars))
}
