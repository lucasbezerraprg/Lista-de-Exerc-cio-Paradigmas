Exercicio 14 em Golang:

package main

import (
    "fmt"
    "sync"
)

type Configuracao struct {
    Parametro string
}

var (
    instance *Configuracao
    once     sync.Once
)

func GetInstance() *Configuracao {
    once.Do(func() {
        instance = &Configuracao{Parametro: "Valor Inicial"}
    })
    return instance
}

func main() {
    config1 := GetInstance()
    fmt.Println(config1.Parametro)  

    config2 := GetInstance()
    config2.Parametro = "Novo Valor"

    fmt.Println(config1.Parametro)  
    fmt.Println(config2.Parametro)  
}