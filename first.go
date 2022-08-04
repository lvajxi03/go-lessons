package main

import( // Groupping imports to be more human-readable
      "fmt"
      "math/rand"
      )

func main() {  // Program entry point
        fmt.Println("Hello, world!")
        fmt.Println("Random number", rand.Intn(100)) // imported "math/rand", however only the last one is used as namespace
}

/*
  Exported names:
  * capital letter in packages' internal identifiers:

  Pi = 3.14 <- exported
  pi = 3.14 <- not exported, internal

*/
