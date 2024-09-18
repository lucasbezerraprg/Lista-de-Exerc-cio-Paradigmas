package main

import "fmt"

type Imprimivel interface {
    Imprimir()
}

type Relatorio struct {
    Titulo string
    Conteudo string
}

func (r Relatorio) Imprimir() {
    fmt.Printf("Relatório: %s\nConteúdo: %s\n", r.Titulo, r.Conteudo)
}

type Contrato struct {
    Nome string
    Termos string
}

func (c Contrato) Imprimir() {
    fmt.Printf("Contrato: %s\nTermos: %s\n", c.Nome, c.Termos)
}

func main() {
    r := Relatorio{Titulo: "Relatório Anual", Conteudo: "Este é o conteúdo do relatório anual."}
    c := Contrato{Nome: "Contrato de Serviço", Termos: "Estes são os termos do contrato de serviço."}

    var i Imprimivel

    i = r
    i.Imprimir()

    i = c
    i.Imprimir()
}
