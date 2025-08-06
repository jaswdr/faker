package faker

import (
	"testing"
)

// Benchmarks for faker.go optimizations

func BenchmarkRandomDigitNot(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.RandomDigitNot(0, 1, 2, 3, 4)
	}
}

func BenchmarkFloat(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Float(2, 1, 100)
	}
}

func BenchmarkFloat32(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Float32(2, 1, 100)
	}
}

func BenchmarkFloat64(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Float64(2, 1, 100)
	}
}

func BenchmarkIntBetweenOptimized(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.IntBetween(1, 1000000)
	}
}

func BenchmarkIntBetweenFullRange(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.IntBetween(-2147483648, 2147483647)
	}
}

// Benchmarks for person.go lazy loading

func BenchmarkPersonFirstName(b *testing.B) {
	f := New()
	p := f.Person()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.FirstName()
	}
}

func BenchmarkPersonLastName(b *testing.B) {
	f := New()
	p := f.Person()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.LastName()
	}
}

func BenchmarkPersonName(b *testing.B) {
	f := New()
	p := f.Person()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Name()
	}
}

// Benchmarks for internet.go optimizations

func BenchmarkInternetIpv6(b *testing.B) {
	f := New()
	internet := f.Internet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = internet.Ipv6()
	}
}

func BenchmarkInternetMacAddress(b *testing.B) {
	f := New()
	internet := f.Internet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = internet.MacAddress()
	}
}

func BenchmarkInternetEmail(b *testing.B) {
	f := New()
	internet := f.Internet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = internet.Email()
	}
}

// Benchmarks for utils.go Shuffle

func BenchmarkShuffleSmallSlice(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(slice))
		copy(tmp, slice)
		_ = Shuffle(tmp)
	}
}

func BenchmarkShuffleLargeSlice(b *testing.B) {
	slice := make([]int, 100)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(slice))
		copy(tmp, slice)
		_ = Shuffle(tmp)
	}
}

// Benchmarks for RandomElementWeighted

func BenchmarkRandomElementWeightedSmall(b *testing.B) {
	f := New()
	elements := map[int]string{
		10: "common",
		5:  "uncommon",
		1:  "rare",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RandomElementWeighted(f, elements)
	}
}

func BenchmarkRandomElementWeightedLarge(b *testing.B) {
	f := New()
	elements := map[int]string{
		1000: "very_common",
		500:  "common",
		100:  "uncommon",
		10:   "rare",
		1:    "very_rare",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RandomElementWeighted(f, elements)
	}
}

// Benchmarks for address.go template processing

func BenchmarkAddressCity(b *testing.B) {
	f := New()
	addr := f.Address()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = addr.City()
	}
}

func BenchmarkAddressStreetName(b *testing.B) {
	f := New()
	addr := f.Address()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = addr.StreetName()
	}
}

func BenchmarkAddressStreetAddress(b *testing.B) {
	f := New()
	addr := f.Address()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = addr.StreetAddress()
	}
}

func BenchmarkAddressFullAddress(b *testing.B) {
	f := New()
	addr := f.Address()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = addr.Address()
	}
}

// Parallel benchmarks to test concurrent performance

func BenchmarkNumerifyParallel(b *testing.B) {
	f := New()
	template := "Order-####-###"
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = f.Numerify(template)
		}
	})
}

func BenchmarkPersonNameParallel(b *testing.B) {
	f := New()
	b.RunParallel(func(pb *testing.PB) {
		p := f.Person()
		for pb.Next() {
			_ = p.Name()
		}
	})
}

func BenchmarkAddressCityParallel(b *testing.B) {
	f := New()
	b.RunParallel(func(pb *testing.PB) {
		addr := f.Address()
		for pb.Next() {
			_ = addr.City()
		}
	})
}

// Memory allocation benchmarks

func BenchmarkNumerifyAllocs(b *testing.B) {
	f := New()
	template := "Order-####-###"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Numerify(template)
	}
}

func BenchmarkIpv6Allocs(b *testing.B) {
	f := New()
	internet := f.Internet()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = internet.Ipv6()
	}
}

func BenchmarkRandomElementWeightedAllocs(b *testing.B) {
	f := New()
	elements := map[int]string{
		1000: "very_common",
		500:  "common",
		100:  "uncommon",
		10:   "rare",
		1:    "very_rare",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RandomElementWeighted(f, elements)
	}
}

func BenchmarkAddressCityAllocs(b *testing.B) {
	f := New()
	addr := f.Address()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = addr.City()
	}
}

// Benchmarks for new optimizations

func BenchmarkRandomStringWithLength(b *testing.B) {
	f := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.RandomStringWithLength(50)
	}
}

func BenchmarkRandomStringWithLengthAllocs(b *testing.B) {
	f := New()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.RandomStringWithLength(50)
	}
}

func BenchmarkLexify(b *testing.B) {
	f := New()
	pattern := "????-????-????"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Lexify(pattern)
	}
}

func BenchmarkLexifyAllocs(b *testing.B) {
	f := New()
	pattern := "????-????-????"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Lexify(pattern)
	}
}

func BenchmarkAsciify(b *testing.B) {
	f := New()
	pattern := "****-****-****"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Asciify(pattern)
	}
}

func BenchmarkAsciifyAllocs(b *testing.B) {
	f := New()
	pattern := "****-****-****"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = f.Asciify(pattern)
	}
}

func BenchmarkPersonNameReplacer(b *testing.B) {
	f := New()
	p := f.Person()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Name()
	}
}

func BenchmarkPersonNameReplacerAllocs(b *testing.B) {
	f := New()
	p := f.Person()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Name()
	}
}

func BenchmarkInternetURLReplacer(b *testing.B) {
	f := New()
	internet := f.Internet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = internet.URL()
	}
}

func BenchmarkInternetURLReplacerAllocs(b *testing.B) {
	f := New()
	internet := f.Internet()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = internet.URL()
	}
}
