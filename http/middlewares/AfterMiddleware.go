package middlewares

import (
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
    "time"
)

func LogLatency(c *gin.Context) {
    t := time.Now()
    c.Next()

    latency := time.Since(t)

    log.WithField("time", latency).Info("route latency")
}
