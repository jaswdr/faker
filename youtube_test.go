package faker

import (
	"fmt"
	"testing"
)

func TestGenerateVideoID(t *testing.T) {
	y := New().YouTube()
	fmt.Println(y)
}
