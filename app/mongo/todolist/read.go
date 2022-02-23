package todolist

import (
	"context"
	"log"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Read(memberName string) []*model.Todolist {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")
	findOption := options.Find()

	var results []*model.Todolist
	filter := bson.D{{Key: "memberName", Value: memberName}}
	cur, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var result model.Todolist
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)

	return results
}
