package handler

import (
	"AA/dto"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
)

type RequestPayload struct {
	AnotherPayLoad string
}

func HanderSlashServer2Wrapper(mongoClient *mongo.Client) func(w http.ResponseWriter, r *http.Request) {

	handle := func(w http.ResponseWriter, r *http.Request) {
		collection := mongoClient.Database("codingchallenge").Collection("amazonstuff3")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Error(err)
			w.Write([]byte("No Data Received,"))
			return
		}

		var v dto.CrawlInfo

		if err := json.Unmarshal(body, &v); err != nil {
			logrus.Errorf("unable to destructure data from server 1 request")
			w.Write([]byte("unable to destructure data from server 1 request"))
			return
		}
		newId, err := collection.InsertOne(context.TODO(), v)
		if err != nil {
			logrus.Errorf("unable to insert to mongo %v\n", err)
			w.Write([]byte("unable to insert to mongo"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Successfully wrote to DB with ID: %s", newId.InsertedID)))

	}

	return handle
}
