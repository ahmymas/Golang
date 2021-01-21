package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i%3 == 0 {
          fmt.Println("fizz")
        }
        if i%5 == 0 {
          fmt.Println("buzz")
        }
        if i%3 != 0 && i%5 != 0 {
           fmt.Println("%d", i)
        }
        fmt.Println("\n")
    }
}
