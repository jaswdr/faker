package faker

import (
	"testing"
)

func TestPokemonJapanese(t *testing.T) {
	v := New().Pokemon().Japanese()
	NotExpect(t, "", v)
	ExpectInString(t, v, pokemonJapanese)
}

func TestPokemonEnglish(t *testing.T) {
	v := New().Pokemon().English()
	NotExpect(t, "", v)
	ExpectInString(t, v, pokemonEnglish)
}

func TestPokemonListLengths(t *testing.T) {
	Expect(t, len(pokemonJapanese), len(pokemonEnglish))
	NotExpect(t, 0, len(pokemonJapanese))
	NotExpect(t, 0, len(pokemonEnglish))
}
