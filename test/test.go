package main

import "fmt"

func main() {

	a := "."
	//b := "+"
	c := "+++.+."

	for i := 0; i < len(c); i++ {

		if string(c[i]) == a {
			fmt.Println("success!")
		}

	}

}
