package member

import (
	"fmt"
	"log"
	"net/http"
	"todolist/app/jwt"
	"todolist/app/mongo/member"
	"todolist/redis"

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
			// create jwt token
			jwtToken, err := jwt.GenerateToken(result.MemberId, result.MemberName)
			if err != nil {
				context.Redirect(http.StatusFound, "/")
				return
			}
			// create redis
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
