package databases

import (
	"fmt"
	"github.com/marcosgfrites/REST-API_Golang/src/utils/common"
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncrypt has responsibility to encrypt the user password
func PasswordEncrypt(userPassword string) (string, error) {
	cost := 8
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), cost)
	if err != nil {
		return "", fmt.Errorf("%s%v", common.PasswordEncryptError, err.Error())
	}

	return string(encryptedPassword), nil
}
