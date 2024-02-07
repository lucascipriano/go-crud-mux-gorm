package handles

import (
	"crud/initializers"
	"crud/models"
	"encoding/json"
	"io"

	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read body"))
		return
	}

	var user models.UserModel

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte("Erro ao convereter para struct"))
		return
	}
	db, err := initializers.ConnectDb()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	w.Write([]byte("User created successfully"))

	result := db.Create(&user)
	if result.Error != nil {
		w.Write([]byte("Failed to create user"))
		return
	}

}

func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	    db, err := initializers.ConnectDb()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Failed to connect to database"))
        return
    }

    var users []models.UserModel
    if err := db.Find(&users).Error; err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Failed to fetch users from database"))
        return
    }

    usersJSON, err := json.Marshal(users)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Failed to convert users to JSON"))
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(usersJSON)
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
