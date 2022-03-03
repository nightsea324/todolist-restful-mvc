package member

import (
	"fmt"
	"log"
	"net/http"
	"todolist/app/mongo/member"
	"todolist/jwt"
	"todolist/redis"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login - 登入會員
func Login(context *gin.Context) {
	var status string
	var msg string

	// 取得登入資訊
	loginName := context.PostForm("loginName")
	loginPassword := context.PostForm("loginPassword")

	// 確認會員資訊是否存在
	result, err := member.GetByName(loginName)
	if err != nil {
		status = "failed"
		msg = "登入失敗，無此名稱"
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(result.MemberPassword), []byte(loginPassword))
		if err != nil {
			status = "failed"
			msg = "登入失敗，密碼錯誤"
			fmt.Println(err)
		} else {
			// 建立jwt
			jwtToken, err := jwt.GenerateToken(result.MemberId, result.MemberName)
			if err != nil {
				context.Redirect(http.StatusFound, "/")
				return
			}
			// 寫入redis
			if err := redis.Set(context, loginName, jwtToken); err != nil {
				log.Fatal(err)
			}
			context.SetCookie(jwt.Key, jwtToken, jwt.JWT_TOKEN_LIFE, "/", "localhost", false, true)

			status = "ok"
			msg = "登入成功"
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": msg,
	})
}
