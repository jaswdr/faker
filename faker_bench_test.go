package faker

import (
	"testing"
)

func BenchmarkNumerify(b *testing.B) {
	f := New()
	template := "Order-####-###"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Numerify(template)
	}
}

func BenchmarkLexify(b *testing.B) {
	f := New()
	template := "User-????-???"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Lexify(template)
	}
}

func BenchmarkBothify(b *testing.B) {
	f := New()
	template := "ID-##??-##??"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Bothify(template)
	}
}

func BenchmarkAsciify(b *testing.B) {
	f := New()
	template := "Key-****-***"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Asciify(template)
	}
}

