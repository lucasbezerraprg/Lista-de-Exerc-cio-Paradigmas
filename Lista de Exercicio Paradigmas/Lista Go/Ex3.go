package main

import (
  "fmt"
)

type ContaBancaria struct {

  titular string
  saldo  float64 
}

func NovaContaBancaria(titular string, saldoInicial float64) *ContaBancaria {

  return &ContaBancaria{

    titular: titular,
    saldo:  saldoInicial,
  }
}

func (c *ContaBancaria) Saldo() float64 {

  return c.saldo
}

func (c *ContaBancaria) Depositar(valor float64) {

  if valor > 0 {
    c.saldo += valor
    fmt.Printf("Depósito de R$ %.2f realizado com sucesso. Saldo atual: R$ %.2f\n", valor, c.saldo)
  } 
  else {
    fmt.Println("Valor de depósito deve ser positivo.")
  }
}

func (c *ContaBancaria) Sacar(valor float64) {

  if valor > 0 {
    if valor <= c.saldo {
      c.saldo -= valor
      fmt.Printf("Saque de R$ %.2f realizado com sucesso. Saldo atual: R$ %.2f\n", valor, c.saldo)
    } 
	else {
      fmt.Println("Saldo insuficiente para saque.")
    }

  } 
  else {
    fmt.Println("Valor de saque deve ser positivo.")
  }
}

func main() {

  conta := NovaContaBancaria("João Silva", 1000.00)
  fmt.Printf("Saldo inicial: R$ %.2f\n", conta.Saldo())
  conta.Depositar(500.00)
  conta.Depositar(-100.00) 
  conta.Sacar(200.00)
  conta.Sacar(1500.00) 
  conta.Sacar(-50.00) 

  fmt.Printf("Saldo final: R$ %.2f\n", conta.Saldo())

}