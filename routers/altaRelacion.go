package routers

import (
	"fmt"
	"net/http"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/bd"
	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
)

/*AltaRelacion relacion el registro de la relacion entre usuarios */
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("llego al servicio   1")
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "el parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se a logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
