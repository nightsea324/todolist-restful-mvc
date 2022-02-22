package route

import (
	"todolist/app/usecase/todolist"

	"github.com/gin-gonic/gin"
)

func Route() {
	router := gin.Default()
	router.POST("/todolist/", todolist.Create)
	router.DELETE("/todolist/", todolist.Delete)
	router.Run(":80")
}
