package Config

import (
    "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "strings"
)

const ENV = "env"
const DATABASE_URL = "database.url"
const PORT = "port"
const ADDRESS = "address"

func defaults() {
    viper.SetEnvPrefix("App")

    viper.SetDefault(ENV, "dev")
    viper.SetDefault(PORT, 8000)
    viper.SetDefault(ADDRESS, "0.0.0.0")
    viper.SetDefault(DATABASE_URL, "mongodb://localhost:27017/my_collection")

}

func Load(force bool) {
    defaults()
    autoload(force)
}

func autoload(force bool) {

    viper.AddConfigPath("./")
    viper.SetConfigName("application")
    if force {
        viper.WriteConfigAs("application.yml");
        logrus.Info("Forcing config to load")
        os.Exit(0)
    }


    if err := viper.ReadInConfig(); err != nil {
        viper.WriteConfigAs("application.yml");
        logrus.Warning("config file generated")
    }
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    viper.AutomaticEnv()
}
