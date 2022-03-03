package todolist

import (
	"net/http"
	"todolist/app/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Update - 完成待辦事項
func Update(context *gin.Context) {

	var status string
	var msg string

	// 確認資料
	todoId := context.Param("id")
	todoList, err := todolist.GetById(todoId)

	if err != nil {
		status = "failed"
		msg = "待辦事項不存在"
	} else {
		// 確認使用者
		if context.GetString("memberName") != todoList.MemberName {
			status = "failed"
			msg = "使用者錯誤"
		} else {
			// 更新資料庫資料
			todolist.Update(todoId)
			status = "ok"
			msg = "已成功完成待辦事項"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}
