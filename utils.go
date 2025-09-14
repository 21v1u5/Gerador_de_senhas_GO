package main

import (
	"fmt"
	"strings"
)

func linha() {
	num_linhas := strings.Repeat("=", 30)
	fmt.Println(num_linhas)
}

func menu() {
	linha()
	fmt.Println("1. Mostrar senhas")
	fmt.Println("2. Criar senha")
	fmt.Println("3. Deletar senha")
	fmt.Println("4. Sair do programa")
	fmt.Print("Qual é a sua escolha? ")
}
