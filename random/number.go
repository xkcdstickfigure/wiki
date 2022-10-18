package random

import (
	"math"
	"math/rand"
	"time"
)

func Number(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9*int(math.Pow(10, float64(length)-1))) + int(math.Pow(10, float64(length)-1))
}
