package main

import (
	"fmt"
	calculateAverage "goroutine1/calculateAverage"
	random "goroutine1/createRandom"
	printaverage "goroutine1/printAverage"
)

func main() {
	chRandInt := make(chan []int)
	chAverage := make(chan int)
	chAverageResult := make(chan int)

	fmt.Println("✅ Створили канали")

	go random.CreateRandom(10, chRandInt)
	numbers := <-chRandInt
	fmt.Println("Отримали рандомні числа з каналу:", numbers)

	go calculateAverage.CalculateAverage(numbers, chAverage)
	average := <-chAverage
	fmt.Println("Отримали середнє число з рандомних чисел з каналу", average)

	go printaverage.PrintAverage(chAverageResult, average)
	result := <-chAverageResult
	fmt.Println("Третя горутина виконана та результат", result)
}
