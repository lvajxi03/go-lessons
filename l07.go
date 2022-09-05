package main

import "fmt"

type MyType struct {
	X, Y int
}

func (mt MyType) print() { // this operates on the data copy
	mt.X += 2
	mt.Y += 2
	fmt.Println(mt.X, mt.Y)
}

func (mt *MyType) print2() { // this can modify the value for future use
	mt.X += 2
	mt.Y += 2
	fmt.Println(mt.X, mt.Y)
}

func main() {
	var m = MyType {X: 2, Y: 3}
	n := &m
	fmt.Println(m)
	// non-pointer operates on copy
	m.print()
	m.print()
	// non-pointer operates on real values
	m.print2()
	m.print2()
	// pointer operates on data copy
	n.print()
	n.print()
	// SIGSEGV
	var p *MyType = nil
	p.print()
}
