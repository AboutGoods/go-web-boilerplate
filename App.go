package main

import (
    "App/database"
    "App/init/Config"
    "App/init/Logger"
    "App/router"
    "App/utils"
    "flag"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "strconv"

    "github.com/spf13/viper"
)

func init() {
    var forceConfig = flag.Bool("reconfigure", false,"override config file ")
    flag.Parse()

    Config.Load(*forceConfig)
    Logger.Load()
    database.NewConnection(viper.GetString(Config.DATABASE_URL))

}

func main() {

    gin.SetMode(gin.ReleaseMode)

    engine := gin.New()

    engine.Use(gin.Recovery())

    router.DeclareRoutes(engine)

    address := viper.GetString(Config.ADDRESS)
    port := viper.GetInt(Config.PORT)

    logrus.WithFields(logrus.Fields{
        "host": address,
        "port": port,
    }).Info("server has started")

    err := engine.Run(address +":"+ strconv.Itoa(port))
    utils.PanicOnError(err)
}
