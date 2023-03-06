package main

import (
	"gin-api/project/db"
	"gin-api/project/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
