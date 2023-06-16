package main

import "fmt"



type flags uint
const(
	FlagUp flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v flags) bool { return v&FlagUp == FlagUp }
func TurnDown(v *flags) { *v &^= FlagUp }
func SetBroadcast(v *flags) { *v |= FlagBroadcast }
func IsCast(v flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }


func main() {
	var v flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}