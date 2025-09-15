package main

import (
	"fmt"
)

func linha() {
	fmt.Println("=======================================================")
}

func titulo() {
	linha()
	fmt.Println("                 Gerador de senhas                 ")
	linha()
}

func menu() {
	fmt.Print("\n")
	titulo()
	fmt.Println("1. Mostrar senhas cadastradas")
	fmt.Println("2. Criar nova senha")
	fmt.Println("3. Deletar senha")
	fmt.Println("4. Sair do programa")
	fmt.Print("Qual é a sua escolha? ")
}
