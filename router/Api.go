package router

import (
    "App/http/controllers"
    "App/http/middlewares"
    "github.com/gin-gonic/gin"
)

func DeclareRoutes(engine *gin.Engine) {

    engine.GET("/ping", controllers.MainController{}.Ping)
    engine.GET("/", controllers.MainController{}.Homepage)

    users := engine.Group("/users")
    users.
        Use(middlewares.LoadUser).
        Use(middlewares.LogLatency)
    {
        userController := controllers.UserController{}
        users.GET("", userController.List)
        users.POST("", userController.Post)
        users.PATCH("/:userId", userController.Patch)
        users.GET("/:userId", userController.Get)
        users.DELETE("/:userId", userController.Delete)
    }
}
