package xrand

import (
	goworktest "github.com/dobyte/go-work-test"
	"math/rand"
)

func Float64() float64 {
	return rand.Float64()
}

func Float32() float32 {
	return rand.Float32()
}

func Int64() int64 {
	return goworktest.Int64()
}

func B() string {
	return "B"
}
