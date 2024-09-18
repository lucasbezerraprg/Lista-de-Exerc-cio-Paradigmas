package main

import (
  "fmt"
)


type Animal interface {
  Som() string
}

type Cachorro struct{}

func (c Cachorro) Som() string {
  return "Au au!"
}

type Gato struct{}

func (g Gato) Som() string {
  return "Miau!"
}

type Vaca struct{}

func (v Vaca) Som() string {
  return "Muu!"
}

func FazerSom(animais []Animal) {
  for _, animal := range animais {
    fmt.Println(animal.Som())
  }
}

func main() {

  cachorro := Cachorro{}
  gato := Gato{}
  vaca := Vaca{}
  animais := []Animal{cachorro, gato, vaca}

  FazerSom(animais)
}