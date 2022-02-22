package todolist

import (
	"context"
	"log"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func Delete(todoId string) {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")
	filter := bson.D{{Key: "todoId", Value: todoId}}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
}
