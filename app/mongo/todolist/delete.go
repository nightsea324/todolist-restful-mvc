package todolist

import (
	"context"
	"log"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete - 刪除待辦事項
func Delete(todoId string) {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	// 刪除資料庫資料
	filter := bson.D{{Key: "todoId", Value: todoId}}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
}
