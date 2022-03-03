package todolist

import (
	"context"
	"log"
	"todolist/app/model"
	"todolist/app/mongo"
)

// Insert - 新增待辦事項
func Insert(todoList model.Todolist) {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	// 寫入資料庫
	_, err := collection.InsertOne(ctx, todoList)
	if err != nil {
		log.Fatal(err)
	}
}
