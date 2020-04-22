package main

import (
	"fmt"
	"math"
)

func main() {
	current, next, sum := 1, 2, 2
	for {
		if next > 4000000 {
			break
		}
		temp := next
		next = current + next
		current = temp
		if IsEven(next) {
			sum += next
		}
	}
	fmt.Println(sum)
}

func IsEven(i int) bool {
	return math.Mod(float64(i), float64(2)) == 0
}
