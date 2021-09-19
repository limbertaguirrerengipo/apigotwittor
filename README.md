





#  heroku 
cd my-project/
$ git init
$ heroku git:remote -a apigo-twittor

Commit your code to the repository and deploy it to Heroku using Git.

$ git add .
$ git commit -am "make it better"
$ git push heroku master

# instalando mongo

$ go get go.mongodb.org/mongo-driver/mongo
$ go get go.mongodb.org/mongo-driver/options
$ go get go.mongodb.org/mongo-driver/bson
$ go get go.mongodb.org/mongo-driver/bson/primitive

$ go get golang.org/x/crypto/bcrypt
$ go get  github.com/gorilla/mux
$ go get github.com/rs/cors
$ go get github.com/dgrijalva/jwt-go

# construir y correr

$ go build main.go
$ go run main.go