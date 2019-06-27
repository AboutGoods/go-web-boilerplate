package controllers

import (
    "App/database"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
}

func (Controller) Collection(collectionName string) *mongo.Collection {
    return database.Collection(collectionName)
}
func (Controller) Log(c *gin.Context) *log.Entry {
    contextualizedLogger := log.WithFields(
        log.Fields{
            "route":c.Request.RequestURI,
        })
    return contextualizedLogger
}
