package todolist

import (
	"net/http"
	"todolist/app/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Read - 查詢待辦事項
func Read(context *gin.Context) {

	var status string
	var msg string

	// 透過使用者查詢
	memberName := context.GetString("memberName")
	results, err := todolist.GetByName(memberName)
	if err != nil {
		status = "failed"
		msg = "使用者無待辦事項"
	} else {
		status = "ok"
		msg = "已查詢使用者待辦事項"
	}

	context.JSON(http.StatusOK, gin.H{
		"result": results,
		"status": status,
		"msg":    msg,
	})
}
