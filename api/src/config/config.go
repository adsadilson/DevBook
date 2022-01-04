package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco responsável para amarzenar a url
	StringConexaoBanco = ""
	// Porta onde a API vai estar rodando
	Porta = 0
)

// Carregar vai iniciar as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Error loading .env file", erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_NOME"),
	)

}
