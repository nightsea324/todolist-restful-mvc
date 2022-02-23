package todolist

import (
	"context"
	"log"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Update(todoId string) {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")
	filter := bson.D{{Key: "todoId", Value: todoId}}
	opts := options.Update().SetUpsert(true)
	update := bson.D{
		{"$set", bson.D{
			{Key: "todoStatus", Value: true}},
		}}
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
}
