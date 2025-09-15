# 🔐 Gerador de Senhas (Golang)

Um gerador de senhas simples feito em **Go**, com opções de criação, visualização e remoção de senhas.

---

## 📌 Funcionalidades

- Escolher o **tamanho da senha**.
- Gerar senha com **letras, números e caracteres especiais**.
- **Mostrar senhas** já criadas.
- **Criar nova senha** → informe o tamanho desejado → senha é armazenada.
- **Deletar senha** → escolha qual senha → remove da lista.
- **Sair do programa** com segurança.

---

## 📂 Estrutura do Projeto

- `listas.go` → listas de caracteres (maiúsculas, minúsculas, números, símbolos).
- `utils.go` → funções básicas de apoio.
- `controler.go` → lógica principal do sistema.
- `main.go` → ponto de entrada do programa.
- `readme.md` → documentação do projeto.

---

## 🚀 Como Rodar o Projeto

1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/gerador-senha-go.git
   cd gerador-senha-go

2. Complile e execute
    ```bash
    go run .
--------------------

🔧 Melhorias Futuras

 Tratamento de erros → (ex.: evitar crash ao digitar letras no menu).

 Persistência de dados → salvar as senhas em arquivo local.

 Testes automatizados → garantir a qualidade do código.