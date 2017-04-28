package main

import (
	"fmt"
	"math"
	"strconv"
)

func DigPow(n, p int) {

	length := len(strconv.Itoa(n))
	var i int
	pwr := p + length - 1
	var sum int
	num := n

	for i < length {

		sum += int(math.Pow(float64(num%10), float64(pwr)))
		num = num / 10
		pwr--
		i++
	}

	if sum%n == 0 {
		fmt.Println(sum / n)
	} else {
		fmt.Println("-1")
	}

}

func main() {
	DigPow(89, 1)    //1
	DigPow(92, 1)    //-1
	DigPow(695, 2)   //2
	DigPow(46288, 3) //51
}
