package todolist

import (
	"net/http"
	"todolist/app/model"
	"todolist/app/mongo/todolist"

	"github.com/gin-gonic/gin"
)

func Read(context *gin.Context) {
	var status string
	var msg string
	memberName := context.Query("memberName")
	var result []*model.Todolist
	result = todolist.Read(memberName)
	status = "ok"
	msg = "已查詢使用者待辦事項"

	context.JSON(http.StatusOK, gin.H{
		"result":  result,
		"status":  status,
		"message": msg,
	})
}
