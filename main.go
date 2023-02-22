package main

import (
	"go-sakila-api/database"
	"go-sakila-api/resources/actors"
	"go-sakila-api/server"
)

func main() {
	database.Init()
	database.DB.AutoMigrate(&actors.Actor{})
	server.Init()
}
