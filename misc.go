package frutils

import (
	"math/rand"
	"time"
)

//GetRandomInt returns an integer between two numbers, inclusively
func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
