package todolist

import (
	"context"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func Check(todoId string) bool {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	var todoList model.Todolist
	filter := bson.D{{Key: "todoId", Value: todoId}}
	err := collection.FindOne(ctx, filter).Decode(&todoList)
	if err != nil {
		return false
	}
	return true
}
