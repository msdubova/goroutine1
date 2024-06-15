package print

import "fmt"

func PrintAverage(chAverageResult chan<- int, average int) {
	fmt.Println("✅ Виконується горутина 3 PrintAverage ")
	chAverageResult <- average
	fmt.Println("Третя горутина виводить число на екран ", average)
}
