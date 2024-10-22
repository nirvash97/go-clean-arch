package client

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://aegisx1:papth0391@experimental-01.8lsgx.mongodb.net/?retryWrites=true&w=majority&appName=Experimental-01"

func InitialServer() (*mongo.Client, error) {
	/// Mongo DB Connect
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	//
	// NOTE ::  Put Defer outside to avoid trigger it when fuuction has been done
	//

	// defer func() {
	// 	if err = client.Disconnect(context.Background()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	var echo bson.M
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&echo); err != nil {
		panic(err)
	}
	return client, err
}
