package Logger

import (
    "App/init/Config"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
)

func Load() {
    log.SetOutput(os.Stdout)
    env := viper.GetString(Config.ENV)
    log.SetLevel(log.DebugLevel)
    if env == "prod" {
        log.SetFormatter(&log.JSONFormatter{})
        log.SetLevel(log.InfoLevel)
    }

}
