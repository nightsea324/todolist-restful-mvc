package route

import (
	"todolist/app/middleware"
	"todolist/app/usecase/member"
	"todolist/app/usecase/todolist"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()
	router.POST("/todolist/", middleware.Auth(), todolist.Create)
	router.DELETE("/todolist/", middleware.Auth(), todolist.Delete)
	router.PUT("/todolist/", middleware.Auth(), todolist.Update)
	router.GET("/todolist/", middleware.Auth(), todolist.Read)
	router.POST("/member/", member.Register)
	router.POST("/member/login", member.Login)
	router.Run(":80")
}
