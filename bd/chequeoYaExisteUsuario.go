package bd

import (
	"context"
	"fmt"
	"time"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un mail de parametro y verificar si existe en la bd  */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("usuarios")
	condicion := bson.M{"Email": email}
	fmt.Println("condicion : ", condicion)
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	fmt.Println(err)
	if err != nil {
		return resultado, true, ID
	}

	return resultado, false, ID
}
