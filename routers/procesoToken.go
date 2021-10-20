package routers

import (
	"errors"
	"strings"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*email valor usado en todos los enponts*/
var Email string

/* IDUsuario es el Id devuelto del modelo, que se usara en todos los endpoint*/
var IDUsuario string

/*ProcesoToken es el proceso par extraer los datos */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrollo_grupoFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Invalido")
	}
	return claims, false, string(""), err
}
