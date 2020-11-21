package requestuuid

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/common/random"
)

// RequestUUID means
func RequestUUID() gin.HandlerFunc {
	return func(c *gin.Context) {

		uuid, err := random.GenUUIDV4()
		if err != nil {
			logrus.WithField("err", err).Panic("middleware requestUuid gen uuid failure")
			c.Abort()
		}

		c.Set("requestUUID", uuid)
		c.Next()
	}
}
