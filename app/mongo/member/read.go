package member

import (
	"context"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

// 確認帳號是否存在
func CheckName(name string) (model.Member, bool) {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	var result model.Member
	filter := bson.D{{Key: "memberName", Value: name}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, false
	}

	return result, true
}
