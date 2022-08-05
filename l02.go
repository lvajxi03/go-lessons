package main

import "fmt"

func function1(x int, y int) int { // Function definition: formal parameters followed by their types, also return type
	return x + y
}

func function2(x, y int) int { // To make definition shorter, type can be specified only once for a group of the same types
	return x * y
}

func swap(x, y int) (int, int) { // Multiple return types
	return y, x
}

func function3(x, y int) (a, b int) { // named return values
    a = x + y
    b = y - x
    return // "aked return"
}

func main() {
    fmt.Println("Function1:", function1(2, 3))
    fmt.Println("Function2:", function2(4, 5))
    var a, b int // Variable definition
    a, b = 7, 9 // Multiple assignments
    fmt.Println("Before swap:", a, b)
    a, b = swap(a, b)
    fmt.Println("After swap:", a, b)
    // Short assignment, without type
    // (compiler will assign the type of initializer
    k := 3
    fmt.Println("This is k:", k)
    /*
    Statement above will cause error, as k is already an int:
    k := "And now for something completely different"
    fmt.Println("This is k also:", k)
    */
    x, y := function3(10, 20)
    fmt.Println("x and y:", x, y)
}
