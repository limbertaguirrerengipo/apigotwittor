package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
	"gitgub.com/limbertaguirrerengipo/apigotwittor/jwt"
	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
)

/*  Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario e y/o Contrasenha invalida"+err.Error(), 400)
		return

	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido"+err.Error(), 400)
		return
	}
	fmt.Println("variable t : ", t.Email, t.Password)
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "Usuario 2 y/o Contrasenha invalida", 400)
		return
	}
	jwtkey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar generar el token"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}
	w.Header().Set("Context-Type", "applition/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
