package member

import (
	"context"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// GetByName - 透過名稱查詢
func GetByName(name string) (model.Member, error) {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	// 查詢資料庫
	var result model.Member
	filter := bson.D{{Key: "memberName", Value: name}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
