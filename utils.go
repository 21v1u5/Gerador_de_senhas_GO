package main

import (
	"bufio"
	"fmt"
	"os"
)

func linha() {
	fmt.Println("=======================================================")
}

func titulo() {
	linha()
	fmt.Println("                 GERADOR DE SENHAS                 ")
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

func limparEntrada() {
	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')
}
