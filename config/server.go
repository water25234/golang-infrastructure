package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	goAddrPort = flag.String("goAddrPort", ":8080", "RESTful api port")
)

var appConfig *AppConfig

func init() {
	godotenv.Load()
}

// AppConfig mean app config struct
type AppConfig struct {
	AppLogPath             string
	DBConnection           string
	DBHost                 string
	DBPort                 int
	DBDatabase             string
	DBUsername             string
	DBPassword             string
	RedisHost              string
	RedisPort              string
	RedisDB                int
	LineChannelSecret      string
	LineChannelAccessToken string
	GoAddrPort             string
}

// SetAppConfig mean set app config
func SetAppConfig() {
	redisDb, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	appConfig = &AppConfig{
		AppLogPath:             os.Getenv("APP_LOG_PATH"),
		DBConnection:           os.Getenv("DB_CONNECTION"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBPort:                 dbPort,
		DBDatabase:             os.Getenv("DB_DATABASE"),
		DBUsername:             os.Getenv("DB_USERNAME"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		RedisHost:              os.Getenv("REDIS_HOST"),
		RedisPort:              os.Getenv("REDIS_PORT"),
		RedisDB:                redisDb,
		LineChannelSecret:      os.Getenv("LINE_CHANNEL_SECRET"),
		LineChannelAccessToken: os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		GoAddrPort:             *goAddrPort,
	}
}

// GetAppConfig mean get app config
func GetAppConfig() *AppConfig {
	return appConfig
}
