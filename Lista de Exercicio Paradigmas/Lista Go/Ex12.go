package main

import "fmt"

type Produto struct {
    Nome  string
    Preco float64
}

func SomarProdutos(p1, p2 Produto) Produto {
    return Produto{
        Nome:  p1.Nome + " & " + p2.Nome,
        Preco: p1.Preco + p2.Preco,
    }
}

func main() {
    produto1 := Produto{Nome: "Produto A", Preco: 50.0}
    produto2 := Produto{Nome: "Produto B", Preco: 30.0}
    produto3 := SomarProdutos(produto1, produto2)

    fmt.Printf("%s: R$%.2f\n", produto3.Nome, produto3.Preco)  
}	