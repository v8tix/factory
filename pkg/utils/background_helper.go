package utils

import (
	"math/rand"
	"time"
)

func computeBinaryInt() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	target := rand.Intn(2) % 2
	return target
}

func SleepRandomTime() {
	time.Sleep(time.Duration(computeBinaryInt()+1) * time.Nanosecond)
}
