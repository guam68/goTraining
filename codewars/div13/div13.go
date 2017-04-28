package main

import (
	"fmt"
	"strconv"
)

func Calc(x int) int {

	digits := make([]int, 0)
	remain := [6]int{1, 10, 9, 12, 3, 4}
	length := len(strconv.Itoa(x))
	num := x
	var sum int
	var counter int

	for length > 0 {

		digits = append(digits, num%10)
		num = num / 10
		length--
	}

	for i := 0; i < len(digits); i++ {
		if i%6 != 0 {
			sum += digits[i] * remain[counter]
		} else {
			counter = 0
			sum += digits[i] * remain[counter]
		}
		counter++
	}

	return sum
}

func Thirt(n int) int {
	stable := Calc(n)
	for stable != Calc(stable) {
		stable = Calc(stable)
	}
	fmt.Print(stable)
	return stable
}

func main() {

	Thirt(1234567)

}
