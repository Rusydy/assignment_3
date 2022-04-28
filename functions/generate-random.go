package functions

import "math/rand"

func RandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}
