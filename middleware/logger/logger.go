package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/water25234/golang-infrastructure/common/random"
	log "github.com/water25234/golang-infrastructure/core/logger"
)

// Logger means
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestUUID := c.MustGet("requestUUID").(string)
		if len(requestUUID) == 0 {
			uuid, err := random.GenUUIDV4()
			if err != nil {
				logrus.WithField("err", err).Panic("middleware logger gen uuid failure")
				c.Abort()
			}
			requestUUID = uuid
		}

		// put here for new.
		log.SetLoggerField(requestUUID)

		c.Next()
	}
}
