package todolist

import (
	"net/http"
	"todolist/app/mongo/todolist"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	var status string
	var msg string
	todoId := context.PostForm("todoId")
	if !todolist.Check(todoId) {
		status = "failed"
		msg = "待辦事項不存在"
	} else {
		todolist.Delete(todoId)
		status = "ok"
		msg = "已成功刪除待辦事項"

	}
	context.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}
