package middlew

import (
	"fmt"
	"net/http"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
)

/*ChequeoBD es el middle que pe permite conocer el estado de la base de datos */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("llega la peticion checkeo ")
		if bd.CheckeoConnection() == 0 {
			http.Error(rw, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
