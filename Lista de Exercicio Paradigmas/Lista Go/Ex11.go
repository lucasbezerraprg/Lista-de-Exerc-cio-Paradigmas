package main

import "fmt"

type Funcionario interface {
    CalcularSalario() float64
}

type FuncionarioHorista struct {
    Nome      string
    HorasTrabalhadas float64
    ValorHora float64
}

func (f FuncionarioHorista) CalcularSalario() float64 {
    return f.HorasTrabalhadas * f.ValorHora
}

type FuncionarioAssalariado struct {
    Nome   string
    SalarioMensal float64
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
    return f.SalarioMensal
}

func main() {
    horista := FuncionarioHorista{Nome: "João", HorasTrabalhadas: 160, ValorHora: 50}
    assalariado := FuncionarioAssalariado{Nome: "Maria", SalarioMensal: 5000}

    var f Funcionario

    f = horista
    fmt.Printf("Salário de %s: %.2f\n", horista.Nome, f.CalcularSalario())

    f = assalariado
    fmt.Printf("Salário de %s: %.2f\n", assalariado.Nome, f.CalcularSalario())
}
