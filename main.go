package main

import (
	"Goroutine1/calculate"
	"Goroutine1/print"
	"Goroutine1/random"
	"fmt"
)

func main() {
	chRandInt := make(chan []int)
	chAverage := make(chan int)
	chAverageResult := make(chan int)

	fmt.Println("Програма відкрита")
	go random.CreateRandom(10, chRandInt)

	go calculate.CalculateAverage(chRandInt, chAverage)

	go print.PrintAverage(chAverage, chAverageResult)
	<-chAverageResult

	fmt.Println("Програма завершена")
}
