package main

import "fmt"

//Not submitted due to broken test. Submitted 28APR17

func mult(num int) {

	var sum int

	for i := 0; i < num; i++ {

		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		} else {
			continue
		}
	}

	fmt.Println(sum)
}

func main() {
	mult(10)
}
