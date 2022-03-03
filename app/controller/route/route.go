package route

import (
	"todolist/app/controller/middleware"
	"todolist/app/usecase/member"
	"todolist/app/usecase/todolist"

	"github.com/gin-gonic/gin"
)

// Route -
func Route() {
	router := gin.Default()
	// 待辦事項
	t := router.Group("api/todolist")
	{
		t.POST("/", middleware.Auth(), todolist.Create)
		t.DELETE("/:id", middleware.Auth(), todolist.Delete)
		t.PUT("/:id", middleware.Auth(), todolist.Update)
		t.GET("/", middleware.Auth(), todolist.Read)
	}
	// 會員
	m := router.Group("api/member")
	{
		m.POST("/register", member.Register)
		m.POST("/login", member.Login)
	}
	router.Run(":80")
}
