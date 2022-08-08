package main

import "fmt"

func main() {
	// Const
	const pi = 4 // World falling apart
	// Basic types:
	var n bool
	n = false
	fmt.Println(n)
	/* Mismatched types below, need to convert types:
	var e float64= 500
	var f uint8 = 3
	fmt.Println(e +f)
	*/
	var st string = "This is a string"
	fmt.Println(st)
	// Complex numbers
	c1 := complex(4, 5)
	c2 := 10 + 11i
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	c3 := c1 + c2
	fmt.Println("sum:", c3)
	// You can only check if two complex numbers are equal:
	if c1 == c2 {
		fmt.Println("c1 == c2")
	} else {
		fmt.Println("c1 != c2")
	}
	/* This one will not work:
	if c1 > c2 {
		fmt.Println("c1 > c2")
	}
	*/
	i := 3
	j := float32(i) // Convert one type to another
	fmt.Println(j)
	/* Also, cannot convert int/float to complex:
	k := complex(j)
	fmt.Println(k)
	*/
	for a := 0; a < 10; a++ {
		c3 *= 2 - 2i
		fmt.Println(c3)
	}
	x, y := 1, 2
	switch {
	case x < 5:
		fmt.Println("x < 5")
	case y > 1:
		fmt.Println("y > 3") // case x < 5 is executed only
	default:
		fmt.Println("default block")
	}

	// Structures
	type mytype struct {
		A, B int
	}

	p := mytype{3, 5}
	p.A = 3
	p.B = 2
	fmt.Println(p)
	r := &p
	r.A += 8
	fmt.Println(r)
	s := &mytype{1, 1}
	fmt.Println(s)
	t := &mytype{A: 2} // B value defaults to 0
	fmt.Println(t)
}
