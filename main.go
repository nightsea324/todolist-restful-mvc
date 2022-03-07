package main

import (
	"todolist/app/controller/router"
	"todolist/app/model/mongo"
	"todolist/app/model/redis"
)

// main -
func main() {
	mongo.Init()
	redis.NewClient()
	router.Router()
}
