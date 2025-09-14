package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var senhas_salvas_slice []string

func main() {
	all_caracteres := append(numeros_slice, letras_slice...)
	all_caracteres = append(all_caracteres, caracteres_slice...)
	linha()
	fmt.Println("      Gerador de senhas      ")

	quebrar_loop := false
	escolha := 0
	for !quebrar_loop {
		rand.Seed(time.Now().UnixNano())
		menu()
		fmt.Scan(&escolha)
		switch escolha {
		case 1: // mostrar
			if len(senhas_salvas_slice) == 0 {
				fmt.Println("\nVocê não tem senhas cadastradas")
			} else {
				fmt.Printf("senhas: %s\n", senhas_salvas_slice)
			}

		case 2: // criar
			caracteres := 0
			fmt.Print("Quantos caracteres você deseja? ")
			fmt.Scan(&caracteres)
			var c strings.Builder
			for x := 0; x < caracteres; x++ {
				sorteio := all_caracteres[rand.Intn(len(all_caracteres))]
				c.WriteString(sorteio)
			}
			senha_temporaria := c.String()
			senhas_salvas_slice = append(senhas_salvas_slice, senha_temporaria)
			fmt.Printf("Senha nova criada, está entre <>: <%s>\n", senha_temporaria)

		case 3: // deletar (ainda falta implementar)
			continue

		case 4: // sair
			quebrar_loop = true
			linha()
			fmt.Println("Saindo do programa...")
		}
	}
}
