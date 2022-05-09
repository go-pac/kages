package utils

import (
	"math/rand"
	"strconv"
)

func RandRange(low, hi int) string {
	return strconv.Itoa(low + rand.Intn(hi-low))
}
