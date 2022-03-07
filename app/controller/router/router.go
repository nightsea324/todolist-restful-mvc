package router

import (
	"todolist/app/controller/middleware"
	"todolist/app/controller/srv/member"
	"todolist/app/controller/srv/todolist"

	"github.com/gin-gonic/gin"
)

// Route -
func Router() {
	router := gin.Default()
	// 待辦事項
	t := router.Group("api/todolist")
	{
		// 新增待辦事項
		t.POST("/", middleware.Auth(), todolist.Create)

		// 刪除待辦事項
		t.DELETE("/:id", middleware.Auth(), todolist.Delete)

		// 完成待辦事項
		t.PUT("/:id", middleware.Auth(), todolist.Update)

		// 取得待辦事項
		t.GET("/", middleware.Auth(), todolist.Read)
	}
	// 會員
	m := router.Group("api/member")
	{
		// 會員註冊
		m.POST("/register", member.Register)

		// 會員登入
		m.POST("/login", member.Login)
	}
	router.Run(":80")
}
