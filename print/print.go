package print

import "fmt"

func PrintAverage(chAverage <-chan int, chAverageResult chan<- int) {
	fmt.Println("✅ Виконується горутина 3 PrintAverage ")
	result := <-chAverage

	fmt.Println("Третя горутина виводить число на екран ", result)
	chAverageResult <- result
	fmt.Println("Результат надруковано, закриваю канал")
	close(chAverageResult)
}
