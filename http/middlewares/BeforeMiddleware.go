package middlewares

import (
    "App/database/models"
    "App/utils"
    "github.com/AboutGoods/go-utils/log"
    "github.com/gin-gonic/gin"
)

func Log(c *gin.Context) {
    log.ResetContext()
    log.AddToContext("path", c.Request.RequestURI)
    log.AddToContext("method", c.Request.Method)
    log.AddToContext("host", c.Request.Host)
    log.Debug("Request received")
}

func LoadUser(c *gin.Context) {

    userId := c.Param("userId")
    if len(userId) > 0 {
        user := models.User{}
        err := user.Load(userId)
        utils.HttpError(c, 404, err)
        c.Set("user", user)
    }
    c.Next()
}
