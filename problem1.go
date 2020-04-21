package main

import (
	"fmt"
	"math"
)

func main() {

	var sum = 0

	for i := 1; i < 1000; i++ {
		f := float64(i)
		if math.Mod(f, float64(3)) == 0 {
			sum += i
		} else if math.Mod(f, float64(5)) == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
