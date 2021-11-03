package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
)

/*ListarUsuarios lista los usuarios */
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enciar el parametro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)
	result, status := bd.LeoUsuariosTodos(IDUsuario, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los datos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
