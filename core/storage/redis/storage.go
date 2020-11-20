package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/water25234/golang-infrastructure/config"
)

type redisServ struct {
	client *redis.Client
}

// New means
func New() Service {
	return &redisServ{}
}

// Init means
func (rs *redisServ) Init() (err error) {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr: config.GetAppConfig().RedisHost + ":" + config.GetAppConfig().RedisPort,
			DB:   config.GetAppConfig().RedisDB,
		})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"pong": pong,
			"err":  err,
		}).Error("failure to connect redis")
		return err
	}

	rs.client = redisClient
	return nil
}

// GetStorage means
func (rs *redisServ) GetStorage() interface{} {
	return rs
}

// Disconnect means
func (rs *redisServ) Disconnect() {
	defer rs.client.Close()
}

// Set means
func (rs *redisServ) Set(key string, value string, t int) (err error) {
	err = rs.client.Set(key, value, time.Duration(t)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get means
func (rs *redisServ) Get(key string) (val string, err error) {
	val, err = rs.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Del means
func (rs *redisServ) Del(key string) (err error) {
	err = rs.client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

// Keys means
func (rs *redisServ) Keys(key string) (val []string, err error) {
	val, err = rs.client.Keys(key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
