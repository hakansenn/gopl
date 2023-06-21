package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var sFlag = flag.String("s", "", "string to hash")
var sha512sum = flag.Bool("sha512sum", false, "use sha512sum format")
var sha256sum = flag.Bool("sha256sum", false, "use sha384sum format")

func main() {
	flag.Parse()

	var comp := "f8c14b2a7024f15661d617ac5533ddacf25f33653bc09773fc7d0fbf99319ab5"
	
	if *sFlag == "" {
		fmt.Println("No string to hash")
		return
	}
	if *sha512sum {
		hash512:= sha512.Sum512([]byte(*sFlag))
		printHash(hash512[:])
		
	} else if *sha256sum {
		hash256:= sha256.Sum256([]byte(*sFlag))
		printHash(hash256[:])
	} 


}

func printHash(hash []byte) {
	for _, b := range hash {
		printByte(b)
	}
	fmt.Println()
}

func printByte(b byte) {
	if b < 16 {
		print("0")
	}
	fmt.Printf("%x", b)
}
