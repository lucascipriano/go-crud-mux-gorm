package handles

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cria usuário"))
}

func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca todos os usuários"))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca um usuário"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza um usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleta usuário"))
}
