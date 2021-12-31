package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registro criando com sucesso!"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar todos os registros!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar registro!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registro autalizado com sucesso!"))
}

func RemoveUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registro removido com sucesso!"))
}
