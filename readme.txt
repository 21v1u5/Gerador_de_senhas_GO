gerador de senha



-voce pode escolher o tamanho da sua senha
ao gerar a senha com letras,numeros e caracteres, voce pode:

    mostrar senhas
    criar senha > quantos caracteres? >> armazenar
    deletar senha > qual senha? > deletar
    sair do programa


pode salvar localmente num arquivo.txt

--------------------

vou fazer o programa em microsservice:

- listas (partes.go)
- funcoes basicas (utils.go)
- funcoes de logica importante para o sistema (controler.go)
- programa principal (main.go)



------------------------


func gerar_slice_de_letras() {
	for i := 'a'; i <= 'z'; i++ {
		letras_slice = append(letras_slice, string(i))
	}
	for b := 'A'; b <= 'Z'; b++ {
		letras_slice = append(letras_slice, string(b))
	}
}

gerar_slice_de_letras()