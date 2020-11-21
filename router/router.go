package router

import (
	"github.com/gin-gonic/gin"

	apiRestfulLiao "github.com/water25234/golang-infrastructure/api/Restful/LiaoLiao"
	apiRestfulShortener "github.com/water25234/golang-infrastructure/api/Restful/shortener"
	rsPostgresql "github.com/water25234/golang-infrastructure/core/register/service/postgresql"
	rsRedis "github.com/water25234/golang-infrastructure/core/register/service/redis"
	rsShortener "github.com/water25234/golang-infrastructure/core/register/service/shortener"
	log "github.com/water25234/golang-infrastructure/middleware/logger"
	requestuuid "github.com/water25234/golang-infrastructure/middleware/requestUuid"
)

// SetupRouter mean setup router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// api
	LiaoLiao := router.Group("/LiaoLiao")
	{
		rsPostgresql.RegisterDBRun()
		LiaoLiao.GET("", apiRestfulLiao.GetLiaoLiaoMessage)
		LiaoLiao.POST("", apiRestfulLiao.PostLiaoLiaoMessage)
	}

	// api
	v1 := router.Group("/api/v1")
	{
		v1.Use(requestuuid.RequestUUID(), log.Logger())
		shortenerRouting := v1.Group("/shortener")
		{
			// put it here for now.
			rsRedis.RegisterRedisRun()
			rsPostgresql.RegisterDBRun()
			rsShortener.RegisterShortenerInterfaceRun()
			shortenerBiz := apiRestfulShortener.Handler(rsShortener.GetShortenerBusiness())

			shortenerRouting.GET("/:shortenerID", shortenerBiz.GetShortenerURL)
		}
	}

	return router
}
