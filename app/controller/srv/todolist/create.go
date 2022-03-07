package todolist

import (
	"net/http"
	"time"
	"todolist/app/model/model"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Create - 新增待辦事項
func Create(context *gin.Context) {

	var status string
	var msg string

	defer func() {
		context.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得資料
	req := new(model.Todolist)
	context.BindJSON(&req)
	memberId := context.GetString("memberId")

	// 寫入model
	todoList := model.Todolist{
		ID:        bson.NewObjectId().Hex(),
		Name:      req.Name,
		Status:    false,
		MemberId:  memberId,
		CreatedAt: time.Now(),
	}

	// 寫入資料庫
	if err := todolist.Create(todoList); err != nil {
		status = "failed"
		msg = "新增失敗，資料庫錯誤"
		return
	}

	status = "ok"
	msg = "已成功新增" + req.Name + "待辦事項"

	return
}
