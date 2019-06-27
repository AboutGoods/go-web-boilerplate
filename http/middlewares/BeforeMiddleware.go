package middlewares

import (
    "App/database/models"
    "App/utils"
    "github.com/gin-gonic/gin"
)

func LoadUser(c *gin.Context) {

    userId := c.Param("userId")
    if len(userId) > 0 {
        user := models.User{}
        err:= user.Load(userId)
        utils.HttpError(c, 404, err)
        c.Set("user", user)
    }
    c.Next()
}
