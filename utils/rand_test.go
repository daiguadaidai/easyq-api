package utils

import (
	"fmt"
	"testing"
)

func TestRandN(t *testing.T) {
	n := RandN(0)
	fmt.Println(n)
}
