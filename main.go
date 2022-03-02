package main

import (
	"todolist/app/mongo"
	"todolist/app/route"
	"todolist/redis"
)

func main() {
	mongo.Init()
	redis.NewClient()
	route.Route()
}
