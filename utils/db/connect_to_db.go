package db

import (
	"context"
	"sync"
	"vacation_requests/utils/rest_errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var clientInstance *mongo.Client
var clientInstanceError rest_errors.RestErr

//Creating one instance of connection singletone
var mongoOnce sync.Once

const (
	CONNECTIONSTRING     = "mongodb+srv://rand:rand@clusteremployeesvacatio.n6aef.mongodb.net/vacation_requests_db?retryWrites=true&w=majority"
	DB_REQUESTS          = "vacation_requests_db"
	COLLECTION_REQUESTS  = "vacation_requests_collection"
	COLLECTION_EMPLOYEES = "employees"
)

//TODO error handling for erros from mongoDB driver
func GetMongoClient() (*mongo.Client, rest_errors.RestErr) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING).SetWriteConcern(writeconcern.New(writeconcern.WMajority()))

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			panic(err)
		}
		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
