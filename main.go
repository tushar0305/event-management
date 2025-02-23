package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/event-management/db"
	"github.com/tushar0305/event-management/routes"
)

func main(){
	db.InitDb()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost
}