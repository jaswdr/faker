package base

import (
	"math/rand"
)

func RandomDigit() int {
	return rand.Uint64() % 10
}
