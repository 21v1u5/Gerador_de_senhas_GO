package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func mostrar(s []string) {
	if len(s) == 0 {
		fmt.Println("\n=> Você não tem senhas cadastradas")
	} else {
		fmt.Print("\n")
		for i, v := range s[0:] {
			fmt.Printf("%d -> %s\n", i+1, v)
		}
		fmt.Print("\n")
	}
}

func criar(s []string) []string {
	caracteres := 0
	fmt.Print("\n=> Quantos caracteres você deseja? ")
	fmt.Scan(&caracteres)
	var c strings.Builder
	for x := 0; x < caracteres; x++ {
		sorteio := all_caracteres[rand.Intn(len(all_caracteres))]
		c.WriteString(sorteio)
	}
	senha_temporaria := c.String()
	senhas_salvas_slice = append(senhas_salvas_slice, senha_temporaria)
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
		for {
			fmt.Print("=> Quer deletar a única senha salva (s/n)? ")
			fmt.Scan(&escolha)
			if escolha == "s" {
				s = []string{}
				fmt.Println(" OK: Senha apagada com sucesso!")
				break
			} else if escolha == "n" {
				fmt.Println("=> Cancelado operação...voltando ao menu")
				break
			} else {
				fmt.Print(" ERROR: Opção inválida...")
			}
		}
		return s
	} else {
		mostrar(s)
		escolha := 0
		x := len(s)
		fmt.Printf("=> Quer deletar qual senha (1 - %d): ", x)
		fmt.Scan(&escolha)

		novo_s := append(s[:(escolha-1)], s[(escolha-1)+1:]...)
		fmt.Printf(" OK: Senha %d apagada com sucesso!\n", escolha)
		mostrar(s)
		return novo_s
	}

}
