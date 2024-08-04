package utils

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

func GetRandomInt(maxValue int) int {
	return rand.Intn(maxValue)
}
