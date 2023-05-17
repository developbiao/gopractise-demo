package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set new seed
	s0 := rand.NewSource(time.Now().UnixNano())
	r0 := rand.New(s0)
	fmt.Print(r0.Intn(100), ",")
	fmt.Print(r0.Intn(100))
	fmt.Println()

	// Not set new seed
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5), +5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()

	// Random float value between nubmers
	rand.Seed(time.Now().UnixNano())
	res := randFloats(1.0, 3.0, 3)
	fmt.Println(res)

	fmt.Println("Main func done!")

}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		// 0.0 <= f < 1.0
		// 0.3 * (30 - 10) = 20 * 0.3
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
