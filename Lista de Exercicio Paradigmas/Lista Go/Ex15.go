package main

import (
    "fmt"
    "errors"
)

type ContaBancaria struct {
    Saldo float64
}

func NovaContaBancaria(saldoInicial float64) *ContaBancaria {
    return &ContaBancaria{Saldo: saldoInicial}
}

func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.Saldo {
        return errors.New("SaldoInsuficienteException: saldo insuficiente para realizar o saque")
    }
    c.Saldo -= valor
    return nil
}

func main() {
    conta := NovaContaBancaria(100.0)

    err := conta.Sacar(50.0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Saque realizado com sucesso. Saldo atual: R$%.2f\n", conta.Saldo)
    }

    err = conta.Sacar(60.0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Saque realizado com sucesso. Saldo atual: R$%.2f\n", conta.Saldo)
    }
}