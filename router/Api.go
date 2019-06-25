package router

import (
    "Antar/controllers"
    "github.com/gin-gonic/gin"
)

func DeclareRoutes(engine *gin.Engine) {

    engine.GET("/ping", controllers.Ping)
    engine.GET("/pong", controllers.Ping)
    engine.GET("/users", controllers.UserList)
    engine.POST("/users", controllers.UserPost)
    engine.DELETE("/users/{userId}", controllers.UserDelete)

}
