package main

import (
  "fmt"
)

type Motor struct {
  Tipo   string
  Potencia int
}

func (m Motor) Info() string {
  return fmt.Sprintf("Tipo: %s, Potência: %d HP", m.Tipo, m.Potencia)
}

type Carro struct {
  Marca string
  Modelo string
  Ano int
  Motor Motor 
}

func (c Carro) Info() string {
  return fmt.Sprintf("Marca: %s, Modelo: %s, Ano: %d, Motor: [%s]", c.Marca, c.Modelo, c.Ano, c.Motor.Info())
}

func main() {
  motor := Motor{
    Tipo: "V8",
    Potencia: 450,
  }

  carro := Carro{
    Marca: "Ford",
    Modelo: "Mustang",
    Ano: 2024,
    Motor: motor,
  }

  fmt.Println(carro.Info())
}
