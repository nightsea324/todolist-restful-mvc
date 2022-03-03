package model

import "time"

// Todolist - 待辦清單資料庫格式
type Todolist struct {
	TodoId     string    `bson:"todoId" json:"todoId"`
	TodoName   string    `bson:"todoName" json:"todoName"`
	TodoStatus bool      `bson:"todoStatus" json:"todoStatus"`
	MemberName string    `bson:"memberName" json:"memberName"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
}
