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

func main() {
  var a Animal
  a = Cachorro{}
  fmt.Printf("Cachorro faz: %s\n", a.Som())
  a = Gato{}
  fmt.Printf("Gato faz: %s\n", a.Som())
}