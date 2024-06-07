package main

import (
	"errors"
	"fmt"
)

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = number + total
	}
	return total
}

func calculate(number1, number2 int, operasi string) (int, error) {
	var (
		result      int
		errorResult error
	)

	switch operasi {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		errorResult = errors.New("Tidak ada Operasi")
	}
	return result, errorResult
}

func main() {
	skor := []int{50, 100, 20, 5}
	total := sum(skor)
	fmt.Println(total)

	fmt.Println("-------------------------")

	result, err := calculate(10, 2, "*")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}
