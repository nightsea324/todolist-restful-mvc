package member

import (
	"net/http"
	"time"
	"todolist/app/model/model"
	"todolist/app/model/mongo/member"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Register - 註冊會員
func Register(context *gin.Context) {

	var status string
	var msg string

	defer func() {
		context.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得註冊資料
	req := new(model.Member)
	context.BindJSON(&req)

	// 確認名稱是否重複
	_, err := member.GetByName(req.Name)
	if err == nil {
		status = "failed"
		msg = "註冊失敗，名稱重複"
		return
	}

	// 密碼加密
	hashPassword, err := hash(req.Password)
	if err != nil {
		status = "failed"
		msg = "註冊失敗，加密密碼錯誤"
		return
	}

	// 寫入model
	data := model.Member{
		ID:        bson.NewObjectId().Hex(),
		Name:      req.Name,
		Password:  hashPassword,
		CreatedAt: time.Now(),
	}

	// 寫入資料庫
	if err := member.Create(data); err != nil {
		status = "failed"
		msg = "註冊失敗，資料庫錯誤"
		return
	}

	status = "ok"
	msg = "註冊成功"

	return
}

// hash -  加密
func hash(password string) (string, error) {

	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
