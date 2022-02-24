package member

import (
	"net/http"
	"todolist/app/mongo/member"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var status string
	var msg string
	memberAccount := context.PostForm("memberAccount")
	memberPassword := context.PostForm("memberPassword")
	if !member.CheckAcount(memberAccount) {
		status = "failed"
		msg = "登入失敗，無此帳號"
	} else {
		if memberPassword != member.GetAccount(memberAccount) {
			status = "failed"
			msg = "登入失敗，密碼錯誤"
		} else {
			status = "ok"
			msg = "登入成功"
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}
