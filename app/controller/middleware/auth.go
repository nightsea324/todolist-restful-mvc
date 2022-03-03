package middleware

import (
	"net/http"
	"todolist/jwt"
	"todolist/redis"

	"github.com/gin-gonic/gin"
)

// Auth - 驗證token
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {

		// 確認cookie
		token, err := context.Cookie(jwt.Key)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "token failed",
				"msg":    "驗證token失敗,無token",
			})
			context.Abort()
			return
		}

		// 解析token
		memberId, memberName, err := jwt.ParseToken(token)
		if err != nil || memberId == "" || memberName == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "token failed",
				"msg":    "驗證token失敗,token錯誤",
			})
			context.Abort()
			return
		}

		// 取得舊token
		oldtoken, err := redis.Get(context, memberName)
		if err != nil || oldtoken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "token failed",
				"msg":    "驗證token失敗,token不存在",
			})
			context.Abort()
			return
		}

		// 驗證token
		if token != oldtoken {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "token failed",
				"msg":    "驗證token失敗,token不符合",
			})
			context.Abort()
			return
		}

		context.Set("memberId", memberId)
		context.Set("memberName", memberName)
		context.Next()
	}
}
