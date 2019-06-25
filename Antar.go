package main

import (
    "Antar/database"
    "Antar/init/Config"
    "Antar/init/Logger"
    "Antar/router"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"

    "github.com/spf13/viper"
)


func init() {
    Config.Load()
    Logger.Load()
    database.NewConnection(viper.GetString(Config.DATABASE_URL))

}


func main() {


    gin.SetMode(gin.ReleaseMode)

    engine := gin.New()
    router.DeclareRoutes(engine)
    logrus.WithFields(logrus.Fields{
        "host": "localhost",
        "port":   8080,
    }).Info("server has started")

    engine.Run()
}
