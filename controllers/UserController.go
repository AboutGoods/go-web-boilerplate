package controllers

import (
    "Antar/database"
    "Antar/database/models"
    "Antar/utils"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func UserList(c *gin.Context) {
    cur, _ := database.Collection(models.USERS).Find(database.Context, bson.D{})
    users := []models.UserApi{}

    if cur != nil {
        defer cur.Close(database.Context)
        for cur.Next(database.Context) {
            var result bson.M
            cur.Decode(&result)
            user := models.UserApi{
                User: models.User{
                    Name:     result["name"].(string),
                    Age:      int(result["age"].(int32)),
                    Company:  result["company"].(string),
                    IsActive: result["isactive"].(bool),
                },
                Id: result["_id"].(primitive.ObjectID).Hex(),
            }
            users = append(users, user)
        }
    }
    c.JSON(200, users)
}

func UserPost(c *gin.Context) {
    var userPost models.UserPost
    err := c.Bind(&userPost)
    utils.PanicOnError(err)
    user := models.User{
        Name:     userPost.Name,
        Company:  userPost.Company,
        Age:      userPost.Age,
        IsActive: true,
    }

    response, err := database.Collection(models.USERS).InsertOne(database.Context, user)
    utils.PanicOnError(err)
    apiUser := models.UserApi{response.InsertedID.(primitive.ObjectID).Hex(), user}
    c.JSON(200, apiUser)
}

func UserGet(c *gin.Context) {
    var userPost models.UserPost
    err := c.BindJSON(&userPost)
    utils.PanicOnError(err)
    user := models.User{
        Name:     userPost.Name,
        Company:  userPost.Company,
        Age:      userPost.Age,
        IsActive: true,
    }

    response, err := database.Collection(models.USERS).InsertOne(database.Context, user)
    utils.PanicOnError(err)
    apiUser := models.UserApi{response.InsertedID.(primitive.ObjectID).String(), user}
    c.JSON(200, apiUser)
}

func UserDelete(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
