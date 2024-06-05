package main

import (
	"fmt"
	"goroutine1/random"
)

func main() {
	numbers := random.CreateRandom(10)
	fmt.Println("Випадкові числа:", numbers)
}
