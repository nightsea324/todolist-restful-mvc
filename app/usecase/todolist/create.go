package todolist

import (
	"net/http"
	"time"
	"todolist/app/model"
	"todolist/app/mongo/todolist"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// Create
func Create(context *gin.Context) {
	var status string
	var msg string
	todoList := model.Todolist{
		TodoId:     bson.NewObjectId().Hex(),
		TodoName:   context.PostForm("todoName"),
		TodoStatus: false,
		MemberName: context.GetString("memberName"),
		CreatedAt:  time.Now(),
	}
	todolist.Insert(todoList)
	status = "ok"
	msg = "新增成功"
	context.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
	})
}
