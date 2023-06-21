package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))

	var s string
	for _, v := range c1 {
		s += fmt.Sprintf("%x ", v)
	}
	fmt.Println(s)

	fmt.Printf("%x \n", c1)

}
