package main

import (
  "fmt"
)
type Carro struct {

  Marca string
  Modelo string
  Ano  int

}

func (c Carro) ExibirDetalhes() {
  fmt.Printf("Marca: %s\nModelo: %s\nAno: %d\n\n", c.Marca, c.Modelo, c.Ano)
}

func main() {
	
  carro1 := Carro{Marca: "Toyota", Modelo: "Corolla", Ano: 2020}
  carro2 := Carro{Marca: "Honda", Modelo: "Civic", Ano: 2021}
  carro3 := Carro{Marca: "Ford", Modelo: "Focus", Ano: 2019}

  carro1.ExibirDetalhes()
  carro2.ExibirDetalhes()
  carro3.ExibirDetalhes()
}