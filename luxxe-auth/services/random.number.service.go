package services

import (
	"math/rand"
	"time"
)

func GetFourRandomNumbers() int {
	seed := time.Now().UnixNano()
	rng  := rand.New(rand.NewSource(seed))
	min  := 1000
	max  := 9999
	return rng.Intn(max-min+1) + min
}
