package calculate

import "fmt"

func CalculateAverage(chRandomInt <-chan []int, chAverage chan<- int) {
	fmt.Println("✅ Виконується горутина 2 calculateAverage")
	numbers := <-chRandomInt
	fmt.Println("Отримали рандомні числа з каналу:", numbers)

	var total int
	for _, item := range numbers {
		total += item
	}

	var average = total / (len(numbers))
	fmt.Println("Середнє число ", average)
	chAverage <- average
	fmt.Println("Середнє число передано в канал, закриваю канал")
	close(chAverage)
}
