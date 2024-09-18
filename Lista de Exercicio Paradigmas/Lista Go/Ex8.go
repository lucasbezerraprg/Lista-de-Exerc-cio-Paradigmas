package main

import (
  "fmt"
)

type Empregado struct {

  Nome  string
  Cargo string
  Salario float64
}

func (e Empregado) Info() string {
  return fmt.Sprintf("Nome: %s, Cargo: %s, Salário: R$ %.2f", e.Nome, e.Cargo, e.Salario)
}

type Empresa struct {
  Nome    string
  Localizacao string
  Empregados []Empregado 
}

func (e *Empresa) AdicionarEmpregado(emp Empregado) {
  e.Empregados = append(e.Empregados, emp)
}

func (e Empresa) Info() string {
  info := fmt.Sprintf("Nome da Empresa: %s, Localização: %s\n", e.Nome, e.Localizacao)
  info += "Empregados:\n"
  for _, emp := range e.Empregados {
    info += emp.Info() + "\n"
  }
  return info
}



func main() {
  emp1 := Empregado{
    Nome:  "Maria Oliveira",
    Cargo:  "Gerente de Projetos",
    Salario: 7500.00,
  }

  emp2 := Empregado{
    Nome:  "João Santos",
    Cargo:  "Desenvolvedor",
    Salario: 5500.00,
  }

  emp3 := Empregado{
    Nome:  "Ana Lima",
    Cargo:  "Analista de Sistemas",
    Salario: 6000.00,
  }

  empresa := Empresa{
    Nome:    "Tech Solutions",
    Localizacao: "São Paulo",
  }

  empresa.AdicionarEmpregado(emp1)
  empresa.AdicionarEmpregado(emp2)
  empresa.AdicionarEmpregado(emp3)
  fmt.Println(empresa.Info())
}