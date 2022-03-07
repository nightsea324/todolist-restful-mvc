package todolist

import (
	"net/http"
	"todolist/app/model/model"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Delete - 刪除待辦事項
func Delete(context *gin.Context) {

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
	req.ID = context.Param("id")
	memberId := context.GetString("memberId")

	// 確認資料是否存在
	todoList, err := todolist.GetById(req.ID)

	if err != nil {
		status = "failed"
		msg = "刪除失敗，待辦事項不存在"
		return
	}

	// 確認使用者是否正確
	if memberId != todoList.MemberId {
		status = "failed"
		msg = "刪除失敗，使用者錯誤"
		return
	}

	// 刪除資料庫資料
	if err := todolist.Delete(req.ID); err != nil {
		status = "failed"
		msg = "刪除失敗，資料庫錯誤"
		return
	}

	status = "ok"
	msg = "已成功刪除待辦事項"

	return
}
