package handlers

import (
	"log"
	"net/http"
	"os"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/middlew"
	"gitgub.com/limbertaguirrerengipo/apigotwittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*MAnejadores seteo mi puerto el habdler m*/
func MAnejadores() {
	router := mux.NewRouter()

	//las rutas
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
