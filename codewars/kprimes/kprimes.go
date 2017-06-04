package main

import (
	"fmt"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		} else {
			continue
		}
	}
	return true
}

/*func isKprime(k int, num int) bool {

	var numPrimes int

	for i := 2; i < num; i++ {
		if num%i == 0 {
			if isPrime(i) == true {
				if i*i == num {
					numPrimes = numPrimes + 2
				} else {
					numPrimes++
				}
			}
		} else {
			continue
		}
	}

	if numPrimes == k {
		return true
	} else {
		return false
	}
}*/

func isKprime(k int, num int) bool {

	var numPrimes int
	knum := num

	for i := 2; knum != 1; i++ {
		if i == num {
			break
		}

		if num%i == 0 {
			if isPrime(i) == true {
				knum = knum / i
				numPrimes++
				fmt.Print(i)
				fmt.Print(" ")
				i = 1
				fmt.Print(numPrimes)
				fmt.Print(" ")
				fmt.Println(knum)
			}
		}
	}
	if numPrimes == k {
		return true
	} else {
		return false
	}

}

func main() {
	//fmt.Println(isKprime(3, 12)) //true
	fmt.Println(isKprime(2, 9)) //true
	//	fmt.Println(isKprime(2, 12)) //false
	//	fmt.Println(isKprime(2, 6))*/ //true
	fmt.Println(isKprime(5, 72)) //true
	//	fmt.Println(isKprime(3, 21))  //false

	//	fmt.Println(isKprime(5, 501)) //false
}
