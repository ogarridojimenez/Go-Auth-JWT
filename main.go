package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ogarridojimenez/go-jwt/controllers"
	"github.com/ogarridojimenez/go-jwt/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
