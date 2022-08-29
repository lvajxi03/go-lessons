package main

import "fmt"


func add(x, y int) int {
    return x + y
}

func operate(x, y int, arg func(a, b int) int) int { // `pointer` to a func
    return arg(x*x, y*y*y)
}

func ret_a_f() func(x, y int) int { // function returning a function
    return func(x, y int) int {
        return x * y
    }
}

func generator() func() int { // closure
    i := 0
    return func() int {
        a := i
        i++
        return a * i
    }
}

func main() {
    x, y := 2, 3
    fmt.Println("add", add(x, y))
    fmt.Println("operate", operate(x, y, add))
    fmt.Println("operate", operate(x + 2, 3 * y, add))
    poi := ret_a_f()
    fmt.Println("poi", poi(2, 3))
    fmt.Println("poi", poi(5, 6))
    gen := generator()
    for i:= 0; i < 10; i++ {
        fmt.Println(gen())
    }
    // this generator is completely separated
    // from the previous one:
    gen2 := generator()
    for i:= 0; i < 10; i++ {
        fmt.Println(gen2())
	}
}
