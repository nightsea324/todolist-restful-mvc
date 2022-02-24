package member

import (
	"context"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckAcount(loginAccount string) (model.Member, bool) {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	var result model.Member
	filter := bson.D{{Key: "memberAccount", Value: loginAccount}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, false
	}

	return result, true
}
