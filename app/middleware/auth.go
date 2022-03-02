package middleware

import (
	"net/http"
	"todolist/app/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := context.Cookie(jwt.Key)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "token failed",
				"msg":    "驗證token失敗,無token",
			})
			context.Abort()
			return
		}

		memberId, memberName, err := jwt.ParseToken(token)
		if err != nil || memberId == "" || memberName == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "token failed",
				"msg":    "驗證token失敗,token錯誤",
			})
			context.Abort()
			return
		}

		context.Set("memberId", memberId)
		context.Set("memberName", memberName)
		context.Next()
	}
}
