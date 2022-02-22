package main

import (
	"todolist/app/mongo"
	"todolist/app/route"
)

func main() {
	mongo.Init()
	route.Route()
}
