package model

import "time"

// Member - 會員資料庫格式
type Member struct {

	// ID - 會員ID
	ID string `bson:"id" json:"id"`

	// Name - 會員名稱
	Name string `bson:"name" json:"name"`

	// Password - 會員密碼
	Password string `bson:"password" json:"password"`

	// CreatedAt - 建立時間
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
