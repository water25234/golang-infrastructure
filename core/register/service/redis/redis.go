package redis

import (
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/core/register"
	storageRedis "github.com/water25234/golang-infrastructure/core/storage/redis"
	rgtrEnum "github.com/water25234/golang-infrastructure/enum/register"
)

// RegisterRedisRun mean
func RegisterRedisRun() {
	register.Run(rgtrEnum.StorageRedis)
}

// GetRedis means
func GetRedis() storageRedis.Service {
	redisDrive, err := register.Get(rgtrEnum.StorageRedis)
	if err != nil {
		logrus.WithField("err", err).Panic("get redis DB failure")
		return nil
	}

	redisService, ok := redisDrive.(*storageRedis.Service)

	if !ok {
		logrus.Panic("trasfer db.Drive to redis.Client is failure")
		return nil
	}
	return *redisService
}
