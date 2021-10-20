package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
)

/*VerPerfil permite ver el perfil de un usuario*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inicio VerPerfil ")
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Deve enviar un parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "pcurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
