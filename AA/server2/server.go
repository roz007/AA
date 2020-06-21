package server1

import (
	"AA/server2/handler"
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRoutes(mongoClient *mongo.Client) *chi.Mux {

	r1 := chi.NewRouter()
	r1.Post("/server2", handler.HanderSlashServer2Wrapper(mongoClient))

	return r1
}

func StartServer2() {
	serverAddr := ":8002"

	cred := options.Credential{
		AuthMechanism:           "SCRAM-SHA-256",
		AuthMechanismProperties: nil,
		AuthSource:              "codingchallenge",
		Username:                "aromal",
		Password:                "bubblegum",
		PasswordSet:             false,
	}
	clientOptions := options.Client().ApplyURI("mongodb://mongo-store:27017").SetAuth(cred)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Error(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Error(err)
	}

	r := SetupRoutes(client)

	err = http.ListenAndServe(serverAddr, r)
	if err != nil {
		logrus.Error(err)
	}

}
