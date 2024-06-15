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

	fmt.Println("✅ Створили канали")

	go random.CreateRandom(10, chRandInt)
	numbers := <-chRandInt
	fmt.Println("Отримали рандомні числа з каналу:", numbers)

	go calculate.CalculateAverage(numbers, chAverage)
	average := <-chAverage
	fmt.Println("Отримали середнє число з рандомних чисел з каналу", average)

	go print.PrintAverage(chAverageResult, average)
	result := <-chAverageResult
	fmt.Println("Третя горутина виконана та результат", result)
}
