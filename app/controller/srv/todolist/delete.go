package todolist

import (
	"net/http"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Delete - 刪除待辦事項
func Delete(context *gin.Context) {

	var status string
	var msg string

	// 確認資料
	id := context.Param("id")
	todoList, err := todolist.GetById(id)

	if err != nil {
		status = "failed"
		msg = "待辦事項不存在"
	} else {
		// 確認使用者
		if context.GetString("memberId") != todoList.MemberId {
			status = "failed"
			msg = "使用者錯誤"
		} else {
			// 刪除資料庫資料
			todolist.Delete(id)
			status = "ok"
			msg = "已成功刪除待辦事項"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}
