package bd

import (
	"context"
	"fmt"
	"time"

	"gitgub.com/limbertaguirrerengipo/apigotwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos lee los usuarios registro en el sistema, si se recibe r en quieres trae solo los que se relacionan comigo */
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twittor")
	col := db.Collection("usuarios")

	var resultado []*models.Usuario

	FindOptions := *options.Find()
	FindOptions.SetLimit(20)
	FindOptions.SetSkip((page - 1) * 20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, &FindOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultado, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true

		}
		if tipo == "follow" && !encontrado {
			incluir = true

		}
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			resultado = append(resultado, &s)
		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	cur.Close(ctx)
	return resultado, true
}
