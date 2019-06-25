package Config

import (
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "strings"
)

const ENV = "env"
const DATABASE_URL = "database.url"


func defaults(){
    viper.SetEnvPrefix("Antar")

    viper.SetDefault(ENV, "dev")
    viper.SetDefault(DATABASE_URL, "mongodb://localhost:27017/my_collection")

}



func Load() {
    defaults()
    autoload()
}


func autoload() {
    viper.AddConfigPath("./")
    viper.SetConfigName("application")
    if err := viper.ReadInConfig(); err != nil {
        viper.WriteConfigAs("application.yml");
        logrus.Info("config file generated")
        os.Exit(0)
    }
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    viper.AutomaticEnv()
}
