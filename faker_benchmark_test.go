package faker

import (
	"testing"
)

func BenchmarkIntBetween(b *testing.B) {
	f := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = f.IntBetween(1, 100)
		}
	})
}
