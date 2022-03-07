package model

import "time"

// Todolist - 待辦事項資料庫格式
type Todolist struct {

	// ID - 待辦事項ID
	ID string `bson:"id" json:"id"`

	// Name - 待辦事項名稱
	Name string `bson:"name" json:"name"`

	// Status - 待辦事項狀態
	Status bool `bson:"status" json:"status"`

	// MemberId - 會員ID
	MemberId string `bson:"memberId" json:"memberId"`

	// CreatedAt - 建立時間
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
