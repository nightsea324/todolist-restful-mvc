package member

import (
	"fmt"
	"net/http"
	"todolist/app/mongo/member"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(context *gin.Context) {
	var status string
	var msg string
	loginName := context.PostForm("loginName")
	loginPassword := context.PostForm("loginPassword")
	result, hasAccount := member.CheckName(loginName)
	if !hasAccount {
		status = "failed"
		msg = "登入失敗，無此名稱"
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(result.MemberPassword), []byte(loginPassword))
		if err != nil {
			status = "failed"
			msg = "登入失敗，密碼錯誤"
			fmt.Println(err)
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
