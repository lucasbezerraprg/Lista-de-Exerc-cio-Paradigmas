package main

import "fmt"

func SomarDois(a, b int) int {
    return a + b
}

func SomarTres(a, b, c int) int {
    return a + b + c
}

func main() {
    resultado1 := SomarDois(3, 4)
    fmt.Printf("A soma de 3 e 4 é: %d\n", resultado1)
	
    resultado2 := SomarTres(1, 2, 3)
    fmt.Printf("A soma de 1, 2 e 3 é: %d\n", resultado2)
}
