package todolist

import (
	"net/http"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Read - 查詢待辦事項
func Read(context *gin.Context) {

	var status string
	var msg string

	// 透過使用者查詢
	id := context.GetString("memberId")
	results, err := todolist.GetByMemberId(id)

	if err != nil {
		status = "failed"
		msg = "查詢失敗，使用者無待辦事項"
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
