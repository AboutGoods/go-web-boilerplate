package utils

import "github.com/gin-gonic/gin"

func HttpError(c *gin.Context, code int, err error) {
    if err != nil {
        c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
    }
}
