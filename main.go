package main

import (
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	factor := NewFactorizer()

	factor.Start()
}
