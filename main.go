package main

import (
	"todolist/app/controller/route"
	"todolist/app/mongo"
	"todolist/redis"
)

// main -
func main() {
	mongo.Init()
	redis.NewClient()
	route.Route()
}
