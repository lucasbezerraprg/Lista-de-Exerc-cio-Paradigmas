package main

import "fmt"

type Matematica struct{}

func (Matematica) Fatorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Matematica{}.Fatorial(n - 1)
}

func main() {
    m := Matematica{}
    fmt.Println(m.Fatorial(5)) 
}