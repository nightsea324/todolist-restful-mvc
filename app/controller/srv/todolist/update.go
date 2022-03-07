package todolist

import (
	"net/http"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Update - 完成待辦事項
func Update(context *gin.Context) {

	var status string
	var msg string

	// 確認資料
	id := context.Param("id")
	todoList, err := todolist.GetById(id)

	if err != nil {
		status = "failed"
		msg = "更新失敗，待辦事項不存在"
	} else {
		// 確認使用者
		if context.GetString("memberId") != todoList.MemberId {
			status = "failed"
			msg = "更新失敗，使用者錯誤"
		} else {
			// 更新資料庫資料
			if err := todolist.Update(id); err != nil {
				status = "failed"
				msg = "更新失敗，資料庫錯誤"
			}
			status = "ok"
			msg = "已成功完成待辦事項"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}
