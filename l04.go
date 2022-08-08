package main

import "fmt"

func main() {
	var a[5]int
	a[2] = 5
	fmt.Println(a)
	p := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(p)
	r := p[2:4]
	fmt.Println(r)
	r[1] = 8
	fmt.Println(p)
	b := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	}
	fmt.Println(b)
	// Slicing a slice
	c := b[1:3]
	fmt.Println(c)
	// Appending
	b = append(b, []int{2, 4, 6, 8, 10})
	fmt.Println(b)
	// Maps
	var m = map[int]string {
		0: "London",
		2: "Bucharest",
		19:"Cairo",
	}
	fmt.Println(m)
	for i, v := range m {
		fmt.Printf("%d => %s\n", i, v)
	}
	delete(m, 2)
	fmt.Println(m)
}
