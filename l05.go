package main

import "fmt"

const (
	A = iota       // 0
	B = iota       // 1
	C = iota       // 2
	D = 5 + iota   // 5 + 3
)


const (
	E = iota       // 0
	F = iota       // 1
	G = iota       // 2
	H = 42 * iota  // 42 * 3
)


type Month int64

const (
	Jan Month = iota  // 0
	Feb               // 1
	Mar               // etc, etc
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

func PrintMonth(m Month) {
	fmt.Println(m)
}

func main() {
	fmt.Println(A, B, C, D)
	fmt.Println(E, F, G, H)
	PrintMonth(Dec)
	PrintMonth(92)
	fmt.Println(A == E)
}

