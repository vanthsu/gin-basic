package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vanthsu/gin-basic/db"
	"github.com/vanthsu/gin-basic/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisteerRoutes(server)

	server.Run(":8080")
}
