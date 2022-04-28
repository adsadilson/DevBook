package services

import (
	"api/src/models"
	"api/src/repository"
)

func Save(usuarioModel models.Usuario) (uint64, error) {

	id, erro := repository.NovaInstanciaUsuario().Save(usuarioModel)

	if erro != nil {
		return 0, erro
	}

	return uint64(id), nil
}
