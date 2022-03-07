package todolist

import (
	"net/http"
	"todolist/app/model/model"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Get - 查詢待辦事項
func Get(context *gin.Context) {

	var status string
	var msg string
	var results []*model.Todolist

	defer func() {
		context.JSON(http.StatusOK, gin.H{
			"results": results,
			"status":  status,
			"message": msg,
		})
	}()

	// 取得使用者資料
	req := new(model.Todolist)
	context.BindJSON(&req)
	memberId := context.GetString("memberId")

	// 透過使用者ID取得資料
	results, err := todolist.GetByMemberId(memberId)

	if err != nil {
		status = "failed"
		msg = "查詢失敗，使用者無待辦事項"
		return
	}

	status = "ok"
	msg = "已查詢使用者待辦事項"

	return
}
