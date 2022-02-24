package route

import (
	"todolist/app/usecase/member"
	"todolist/app/usecase/todolist"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()
	router.POST("/todolist/", todolist.Create)
	router.DELETE("/todolist/", todolist.Delete)
	router.PUT("/todolist/", todolist.Update)
	router.GET("/todolist/", todolist.Read)
	router.POST("/todolist/member", member.Register)
	router.Run(":80")
}
