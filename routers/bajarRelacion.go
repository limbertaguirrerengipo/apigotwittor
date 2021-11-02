package routers

import (
	"net/http"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
)

/* */
func BajarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio al borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se a logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
