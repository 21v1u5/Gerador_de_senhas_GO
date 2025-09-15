package main

import (
	"fmt"
	"math/rand"
	"time"
)

var senhas_salvas_slice []string

func main() {

	quebrar_loop := false
	escolha := 0
	for !quebrar_loop {
		rand.Seed(time.Now().UnixNano())

		menu()
		fmt.Scan(&escolha)

		switch escolha {
		case 1: // mostrar
			mostrar(senhas_salvas_slice)

		case 2: // criar
			criar(senhas_salvas_slice)

		case 3: // deletar (ainda falta implementar)
			senhas_salvas_slice = deletar(senhas_salvas_slice)

		case 4: // sair
			quebrar_loop = true
			linha()
			fmt.Println("Saindo do programa...")
		}
	}
}
