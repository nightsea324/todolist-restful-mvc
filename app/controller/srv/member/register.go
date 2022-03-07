package member

import (
	"log"
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

	data := model.Member{
		ID:        bson.NewObjectId().Hex(),
		Name:      context.PostForm("name"),
		Password:  hash(context.PostForm("password")),
		CreatedAt: time.Now(),
	}

	// 確認名稱是否重複
	_, err := member.GetByName(data.Name)
	if err == nil {
		status = "failed"
		msg = "註冊失敗，名稱重複"
	} else {
		// 寫入資料庫
		member.Create(data)
		status = "ok"
		msg = "註冊成功"

	}

	context.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
	})
}

// hash -  加密
func hash(password string) string {

	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
