package controllers

import (
    "App/database/models"
    "App/utils"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    Controller
}

func (controller UserController) List(c *gin.Context) {
    users := models.UserRepository{}.FindAll()
    c.JSON(200, users)
}

func (controller UserController) Post(c *gin.Context) {
    var user models.User
    err := c.Bind(&user)
    utils.HttpError(c, 400, err)
    user.IsActive = true
    user.Store()
    utils.HttpError(c, 500, err)
    c.JSON(200, user)
}

func (controller UserController) Patch(c *gin.Context) {
    if value, ok := c.Get("user"); ok {
        user := value.(models.User)
        var userPatch models.UserPatch
        err := c.Bind(&userPatch)
        utils.HttpError(c, 400, err)
        user.ApplyPatch(userPatch)
        user.Store()
        c.JSON(202, user)
    }

}

func (controller UserController) Get(c *gin.Context) {
    if value, ok := c.Get("user"); ok {
        user := value.(models.User)
        c.JSON(200, user)
    }
}

func (controller UserController) Delete(c *gin.Context) {
    if value, ok := c.Get("user"); ok {
        user := value.(models.User)
        err := user.Drop()
        utils.HttpError(c, 500, err)
        c.JSON(204, user)
    }
}
