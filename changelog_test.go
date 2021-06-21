package changelog

import (
	"fmt"
	"testing"
)

func TestGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Changelog")
	}
}
