package todolist

import (
	"context"
	"log"
	"todolist/app/model"
	"todolist/app/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetByName - 透過名稱查詢
func GetByName(Name string) ([]*model.Todolist, error) {

	var ctx = context.TODO()

	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	// 查詢資料庫
	findOption := options.Find()
	var results []*model.Todolist
	filter := bson.D{{Key: "memberName", Value: Name}}
	cur, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	cur.Close(ctx)

	return results, nil
}

// GetById - 透過ID查詢
func GetById(id string) (*model.Todolist, error) {
	var ctx = context.TODO()
	// 連線至collection
	collection := mongo.Client.Database("todoList").Collection("todoList")

	result := new(model.Todolist)

	filter := bson.D{{Key: "todoId", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
