package main

import (
  "fmt"
)

type Professor struct {
  Nome   string
  Disciplinas []string
  Escolas []*Escola 
}

type Escola struct {
  Nome    string
  Localizacao string
  Professores []*Professor 
}

func (e *Escola) AdicionarProfessor(p *Professor) {
  e.Professores = append(e.Professores, p)
  found := false

  for _, escola := range p.Escolas {
    if escola == e {
      found = true
      break
    }
  }

  if !found {
    p.Escolas = append(p.Escolas, e)
  }
}

func (p Professor) Info() string {
  return fmt.Sprintf("Nome: %s, Disciplinas: %v", p.Nome, p.Disciplinas)
}

func (e Escola) Info() string {
  return fmt.Sprintf("Nome: %s, Localização: %s", e.Nome, e.Localizacao)
}

func main() {

  escola1 := Escola{
    Nome:    "Escola Primária",
    Localizacao: "Centro da Cidade",
  }

  escola2 := Escola{
    Nome:    "Escola Secundária",
    Localizacao: "Bairro Norte",
  }

  professor1 := Professor{
    Nome:   "Ana Souza",
    Disciplinas: []string{"Matemática", "Física"},
  }

  professor2 := Professor{
    Nome:   "Carlos Pereira",
    Disciplinas: []string{"História", "Geografia"},
  }

  escola1.AdicionarProfessor(&professor1)
  escola1.AdicionarProfessor(&professor2)
  escola2.AdicionarProfessor(&professor1)
  fmt.Println("Informações das Escolas:")
  fmt.Println(escola1.Info())
  fmt.Println("Professores:")

  for _, prof := range escola1.Professores {
    fmt.Println(prof.Info())
  }

  fmt.Println("\nInformações da Segunda Escola:")
  fmt.Println(escola2.Info())
  fmt.Println("Professores:")

  for _, prof := range escola2.Professores {
    fmt.Println(prof.Info())
  }

  fmt.Println("\nEscolas onde o Professor Ana Souza leciona:")
  for _, escola := range professor1.Escolas {
    fmt.Println(escola.Info())
  }
}