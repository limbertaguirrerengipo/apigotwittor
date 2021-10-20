package middlew

import (
	"net/http"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Error en el token ! "+err.Error(), http.StatusBadRequest)
		return
	}
	next.ServeHTTP(w, r)
}
