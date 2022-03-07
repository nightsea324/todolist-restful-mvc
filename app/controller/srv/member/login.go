package member

import (
	"net/http"
	"todolist/app/model/model"
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

	defer func() {
		context.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得登入資訊
	req := new(model.Member)
	context.BindJSON(&req)

	// 驗證會員
	result, err := member.GetByName(req.Name)
	if err != nil {
		status = "failed"
		msg = "登入失敗，無此名稱"
		return
	}

	// 加密密碼比對
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		status = "failed"
		msg = "登入失敗，密碼錯誤"
		return
	}

	// 建立jwt
	jwtToken, err := jwt.GenerateToken(result.ID, result.Name)
	if err != nil {
		context.Redirect(http.StatusFound, "/")
		status = "failed"
		msg = "登入失敗，jwt建立失敗"
		return
	}

	// 寫入redis
	if err := redis.Set(context, req.Name, jwtToken); err != nil {
		status = "failed"
		msg = "登入失敗，redis寫入失敗"
		return
	}

	// 設置cookie
	context.SetCookie(jwt.Key, jwtToken, jwt.JWT_TOKEN_LIFE, "/", "localhost", false, true)

	status = "ok"
	msg = "登入成功"

	return
}
