package member

import (
	"context"
	"log"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckAcount(memberAccount string) bool {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	var result model.Member
	filter := bson.D{{Key: "memberAccount", Value: memberAccount}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false
	}
	return true
}

func GetAccount(memberAccount string) string {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("member")

	var result model.Member
	filter := bson.D{{Key: "memberAccount", Value: memberAccount}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result.MemberPassword
}
