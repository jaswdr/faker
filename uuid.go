package faker

import (
	"crypto/rand"
	"fmt"
)

// UUID is a faker struct for UUID
type UUID struct {
	Faker *Faker
}

// V4 returns a fake UUID version 4
func (UUID) V4() (uuid string) {
	var uiq [16]byte
	rand.Read(uiq[:])
	uiq[6] = (uiq[6] & 0x0f) | 0x40 // Version 4
	uiq[8] = (uiq[8] & 0x3f) | 0x80 // Variant RFC4122
	return fmt.Sprintf("%x-%x-%x-%x-%x", uiq[0:4], uiq[4:6], uiq[6:8], uiq[8:10], uiq[10:])
}
