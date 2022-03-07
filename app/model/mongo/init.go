package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Init - 初始化mongo
func Init() {

	var ctx = context.TODO()

	// 連線至MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongo1:27017,mongo2:27027,mongo3:27037/?replicaSet=rs0")
	mclient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 確認連線
	err = mclient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Client = mclient
}
