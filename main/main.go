package main

import (
	"math/rand"
	"time"
)

func main() {
}

func GenerateSlice(size int) []int {
	s := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return s
}
