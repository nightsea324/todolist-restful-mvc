package todolist

import (
	"context"
	"todolist/app/model/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// Delete - 刪除待辦事項
func Delete(id string) error {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	// 刪除資料庫資料
	filter := bson.D{{Key: "id", Value: id}}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
