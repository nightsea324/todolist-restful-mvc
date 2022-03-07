package member

import (
	"context"
	"todolist/app/model/model"
	"todolist/app/model/mongo"
)

// Create - 註冊會員
func Create(member model.Member) error {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	// 寫入資料庫
	_, err := collection.InsertOne(ctx, member)
	if err != nil {
		return err
	}
	return nil
}
