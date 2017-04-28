package main

import (
	"fmt"
	"math/big"
)

/* math/big package used to handle the large number (max for 64bit is 2^64).
SetString takes the binary string and sets it to the variable "msg" with
the specified base (2). Bytes() used to break up the number into bytes.
The loop then outputs each byte formatted to utf-8.
*/

func main() {
	msg := new(big.Int)
	msg.SetString("0111010001110010011101010111001101110100010101000100100000110001", 2)
	t := msg.Bytes()

	for _, i := range t {
		fmt.Printf("%c", i)
	}
}
