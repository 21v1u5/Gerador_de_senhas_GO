package main

import (
	"fmt"
	"math/rand"
	"time"
)

var senhas_salvas_slice []string

func main() {
	rand.Seed(time.Now().UnixNano())

	quebrar_loop := false
	escolha := 0

	for !quebrar_loop {
		menu()

		if _, err := fmt.Scan(&escolha); err != nil {
			fmt.Println(" ERROR: Entrada inválida. Tente novamente.")
			limparEntrada()
			continue
		}

		switch escolha {
		case 1:
			mostrar(senhas_salvas_slice)
		case 2:
			senhas_salvas_slice = criar(senhas_salvas_slice)
		case 3:
			senhas_salvas_slice = deletar(senhas_salvas_slice)
		case 4:
			linha()
			fmt.Println("\n => Saindo do programa...")
			quebrar_loop = true
		default:
			fmt.Println(" ERROR: Opção inválida.")
		}
	}
}
