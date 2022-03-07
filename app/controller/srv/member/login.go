package member

import (
	"fmt"
	"log"
	"net/http"
	"todolist/app/model/mongo/member"
	"todolist/app/model/redis"
	"todolist/jwt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login - 登入會員
func Login(context *gin.Context) {
	var status string
	var msg string

	// 取得登入資訊
	name := context.PostForm("name")
	password := context.PostForm("password")

	// 確認會員資訊是否存在
	result, err := member.GetByName(name)

	if err != nil {
		status = "failed"
		msg = "登入失敗，無此名稱"
	} else {
		// 加密密碼比對
		err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
		if err != nil {
			status = "failed"
			msg = "登入失敗，密碼錯誤"
			fmt.Println(err)
		} else {
			// 建立jwt
			jwtToken, err := jwt.GenerateToken(result.ID, result.Name)
			if err != nil {
				context.Redirect(http.StatusFound, "/")
				return
			}

			// 寫入redis
			if err := redis.Set(context, name, jwtToken); err != nil {
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
