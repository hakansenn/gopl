package main

import "fmt"

func main() {
	s := "Şen"
	//s := "Sen"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
