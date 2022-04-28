package repository

import (
	"api/src/database"
	"api/src/models"
	"database/sql"
	"fmt"
	"log"
)

type usuarioRepository struct {
	conn *sql.DB
}

// NovaInstanciaUsuario criar uma instancia de usuarioRepository
func NovaInstanciaUsuario() *usuarioRepository {

	conn, erro := database.Conectar()

	if erro != nil {
		log.Fatal(erro)
	}

	return &usuarioRepository{conn}
}

// Save persistir um usuário no banco de dados
func (user usuarioRepository) Save(usuario models.Usuario) (int64, error) {

	query := "insert into usuario (nome, nick, email, senha, status) values (?, ?, ?, ?, ?)"

	statement, erro := user.conn.Prepare(query)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha, usuario.Status)

	if erro != nil {
		return 0, erro
	}

	id, _ := resultado.LastInsertId()

	rowsAffected, _ := resultado.RowsAffected()

	fmt.Print(rowsAffected)

	return id, nil

}
