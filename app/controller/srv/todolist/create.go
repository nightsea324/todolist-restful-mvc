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

	// 寫入資料庫
	todoList := model.Todolist{
		ID:        bson.NewObjectId().Hex(),
		Name:      context.PostForm("name"),
		Status:    false,
		MemberId:  context.GetString("memberId"),
		CreatedAt: time.Now(),
	}

	todolist.Create(todoList)

	status = "ok"
	msg = "新增成功"

	context.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
	})
}
