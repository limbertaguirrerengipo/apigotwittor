package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoC es el objeto de conexion ala base de datos */
var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://laguirre:men75662710@cluster0.3jxjb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*conectarBD es la funcion permite conectar a la bd*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexion exitosa con la bd")
	return client
}

/*CheckeoConnection le hace ping a la base de datos */
func CheckeoConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
