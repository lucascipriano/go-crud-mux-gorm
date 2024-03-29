package handles

import (
	"crud/initializers"
	"crud/models"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	params := mux.Vars(r)
	ID := params["id"]

	db, err := initializers.ConnectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to connect to database"))
		return
	}

	var user models.UserModel
	if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to convert user to JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	db, err := initializers.ConnectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to connect to database"))
		return
	}
	var existingUser models.UserModel
	if err := db.First(&existingUser, ID).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	var updatedUser models.UserModel
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body"))
		return
	}

	existingUser.Name = updatedUser.Name
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password
	existingUser.UpdatedAt = time.Now()

	if err := db.Save(&existingUser).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to update user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	db, err := initializers.ConnectDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to connect to database"))
		return
	}

	var user models.UserModel
	if err := db.Where("id = ?", ID).Delete(&user).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	w.Write([]byte("Deletado"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
