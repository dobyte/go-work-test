package xrand

import (
	goworktest "github.com/dobyte/go-work-test"
	"math/rand"
)

func Int() int {
	return rand.Int()
}

func Int32() int32 {
	return rand.Int31()
}

func Int64() int64 {
	return goworktest.Int64()
}
