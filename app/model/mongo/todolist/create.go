package todolist

import (
	"context"
	"todolist/app/model/model"
	"todolist/app/model/mongo"
)

// Create - 新增待辦事項
func Create(todoList model.Todolist) error {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	// 寫入資料庫
	_, err := collection.InsertOne(ctx, todoList)
	if err != nil {
		return err
	}

	return nil
}
