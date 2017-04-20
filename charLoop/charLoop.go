package main

import "fmt"

func main() {
/* fmt package used to print first 200 characters in a loop. Format 'verbs' are used to modify
	the output text. See: string formatting	 
*/
	for i:=0; i<200; i++{
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
