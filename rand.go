package goworktest

import "math/rand"

func Int64() int64 {
	return rand.Int63()
}

func C() string {
	return "C"
}
