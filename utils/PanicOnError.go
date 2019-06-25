package utils

import (
    log "github.com/sirupsen/logrus"
)

func PanicOnError(err error){
    if err != nil {
        log.Fatal(err)

    }
}