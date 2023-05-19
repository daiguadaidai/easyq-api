package utils

import (
	"math/rand"
	"time"
)

func RandN(n int) int {
	if n == 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandNInt64(n int64) int64 {
	return int64(RandN(int(n)))
}
