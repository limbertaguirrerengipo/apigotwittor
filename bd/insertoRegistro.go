package bd

import (
	"context"
	"time"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es el ue inserta en la base de datos */
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
