package routes

import "github.com/gin-gonic/gin"

func RegisteerRoutes(server *gin.Engine) {
	server.GET("/user", getUsers) // 測試規格是用 user 取清單，跟Laravel的慣例用 users 有差異
	server.GET("/user/:id", getUser)
	server.POST("/user", createUser)
	server.PUT("/user/:id", updateUser)
	server.DELETE("/user/:id", deleteUser)
}
