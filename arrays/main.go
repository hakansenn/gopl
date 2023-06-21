package main

import "fmt"

func main() {
	a:= [2]int{1,2}
	b:= [...]int{1,2}
	c := [2]int{1,3}

	fmt.Println(a == b, a == c, b == c)


}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

