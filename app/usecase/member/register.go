package member

import (
	"log"
	"net/http"
	"time"
	"todolist/app/model"
	"todolist/app/mongo/member"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Register
func Register(context *gin.Context) {
	var status string
	var msg string
	memberdata := model.Member{
		MemberId:       bson.NewObjectId().Hex(),
		MemberName:     context.PostForm("memberName"),
		MemberAccount:  context.PostForm("memberAccount"),
		MemberPassword: hash(context.PostForm("memberPassword")),
		CreatedAt:      time.Now(),
	}
	member.Insert(memberdata)
	status = "ok"
	msg = "註冊成功"
	context.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
	})
}

// 加密
func hash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
