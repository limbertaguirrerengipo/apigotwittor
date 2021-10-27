package bd

import (
	"fmt"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*intentoLogin realiza el chequeo */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	fmt.Println("encontrado 3 : ", encontrado)
	if !encontrado {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	fmt.Println("pass 1 : ", password)
	fmt.Println("pass 2: ", usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
