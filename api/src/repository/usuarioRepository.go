package repository

import (
	"api/src/database"
	"api/src/models"
	"database/sql"
	"log"
)

type usuarioRepository struct {
	con *sql.DB
}

// NovaInstanciaUsuario criar uma instancia de usuarioRepository
func NovaInstanciaUsuario() *usuarioRepository {
	con, erro := database.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	return &usuarioRepository{con}
}

// Save insere um usuário no bnaco de dados
func (user usuarioRepository) Save(usuario models.Usuario) (uint64, error) {
	query := "insert into usuario (nome, nick, email, senha, status) values (?, ?, ?, ?, ?)"
	statement, erro := user.con.Prepare(query)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha, usuario.Status)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}
