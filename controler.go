package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func mostrar(s []string) {
	if len(s) == 0 {
		fmt.Println("\n=> Nenhuma senha cadastrada.")
	} else {
		fmt.Print("\n")
		for i, v := range s[0:] {
			fmt.Printf("%d -> %s\n", i+1, v)
		}
		fmt.Print("\n")
	}
}

func criar(s []string) []string {
	var tamanho int
	fmt.Print("\n=> Qual o tamanho da senha você deseja? ")

	if _, err := fmt.Scan(&tamanho); err != nil || tamanho <= 0 {
		fmt.Println(" ERROR: Tamanho inválido.")
		limparEntrada()
		return s
	}

	var c strings.Builder
	for x := 0; x < tamanho; x++ {
		sorteio := all_caracteres[rand.Intn(len(all_caracteres))]
		c.WriteString(sorteio)
	}
	senha_temporaria := c.String()
	s = append(s, senha_temporaria)
	fmt.Printf(" OK: Senha <%s> criada com sucesso!\n", senha_temporaria)
	return s
}

func deletar(s []string) []string {
	if len(s) == 0 {
		fmt.Println("\n=> Você nao tem senhas para deletar. Voltando ao menu...")
		return s
	} else if len(s) == 1 {
		mostrar(s)
		escolha := ""

		fmt.Print("\n=> Quer deletar a única senha salva (s/n)? ")

		if _, err := fmt.Scan(&escolha); err != nil {
			fmt.Println(" ERROR: Entrada inválida...")
			limparEntrada()
		}

		switch escolha {
		case "s":
			s = []string{}
			fmt.Println(" OK: Senha apagada com sucesso!")
		case "n":
			fmt.Println("=> Cancelado operação...voltando ao menu")
		default:
			fmt.Print(" ERROR: Opção inválida...")
		}

		return s
	} else {
		mostrar(s)
		var escolha int
		x := len(s)
		fmt.Printf("=> Quer deletar qual senha (1 - %d): ", x)

		if _, err := fmt.Scan(&escolha); err != nil || escolha <= 0 || escolha > x {
			fmt.Println(" ERROR: Indice inválido.")
			limparEntrada()
			return s
		}

		novo_s := append(s[:(escolha-1)], s[(escolha-1)+1:]...)
		fmt.Printf(" OK: Senha %d apagada com sucesso!\n", escolha)
		mostrar(novo_s)
		return novo_s
	}

}
