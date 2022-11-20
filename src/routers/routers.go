package routers

import (
	"encoding/json"
	"github.com/marcosgfrites/REST-API_Golang/src/models"
	"github.com/marcosgfrites/REST-API_Golang/src/utils/common"
	"net/http"
)

// Register allows to create a new user in the database
func Register(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, common.RegisterError+err.Error(), http.StatusBadRequest)
		return
	}

	// validate if the email isn't empty because is required
	if len(newUser.Email) == 0 {
		http.Error(w, common.RequiredEmailError, http.StatusBadRequest)
		return
	}

	// validate if the password length is greater than 8 characters
	if len(newUser.Password) < 8 {
		http.Error(w, common.PasswordLengthError, http.StatusBadRequest)
		return
	}

	_, foundedUser, _ := database.UserAlreadyExists(newUser.Email)
	if foundedUser == true {
		http.Error(w, common.UserAlreadyExists, http.StatusBadRequest)
		return
	}

	_, status, err := database.UserSave(newUser)
	if err != nil {
		http.Error(w, common.UserSaveError+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, common.UserPersistError, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
