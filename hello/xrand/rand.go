package xrand

import "math/rand"

func Int() int {
	return rand.Int()
}

func Int32() int32 {
	return rand.Int31()
}
