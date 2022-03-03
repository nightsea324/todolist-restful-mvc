package member

import (
	"context"
	"log"
	"todolist/app/model"
	"todolist/app/mongo"
)

// Insert - 註冊會員
func Insert(member model.Member) {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	// 寫入資料庫
	_, err := collection.InsertOne(ctx, member)
	if err != nil {
		log.Fatal(err)
	}
}
