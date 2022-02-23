package todolist

import (
	"context"
	"log"
	"todolist/app/model"
	"todolist/app/mongo"
)

func Insert(todoList model.Todolist) {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	_, err := collection.InsertOne(ctx, todoList)
	if err != nil {
		log.Fatal(err)
	}
}
