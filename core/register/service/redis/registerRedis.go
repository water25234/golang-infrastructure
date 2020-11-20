package redis

import (
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/core/register"
	storageRedis "github.com/water25234/golang-infrastructure/core/storage/redis"
	rgtrEnum "github.com/water25234/golang-infrastructure/enum/register"
)

// RedisService mean
type RedisService struct {
	redisServ *storageRedis.Service
}

func init() {
	register.Register(rgtrEnum.StorageRedis, &RedisService{})
}

// Run means
func (impl *RedisService) Run() (err error) {
	registerNew := storageRedis.New()

	err = registerNew.Init()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("register redis execute run is failure")
		return err
	}

	redisService := registerNew.GetStorage().(storageRedis.Service)
	impl.redisServ = &redisService
	return nil
}

// Get means
func (impl *RedisService) Get() interface{} {
	return impl.redisServ
}
