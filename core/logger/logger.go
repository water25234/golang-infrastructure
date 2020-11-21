package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/water25234/golang-infrastructure/common/random"
)

// Record means
var Record *loggerField

var log *logrus.Logger

type loggerField struct {
	logrus.FieldLogger
}

// SetLoggerConfig means
func SetLoggerConfig() {
	log = logrus.New()

	if len(os.Getenv("LOG_SERVER")) > 0 {
		log.SetOutput(os.Stderr)
	} else {
		log.SetFormatter(&logrus.JSONFormatter{})

		appLogPath := os.Getenv("APP_LOG_PATH")
		if !fileExists(appLogPath) {
			path := strings.Split(appLogPath, "/")
			join := strings.Join(path[:len(path)-1], "/")
			os.Mkdir(join, 0755)
		}

		file, err := os.OpenFile(appLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	}
}

// SetLoggerField means
func SetLoggerField(uuid string) {
	if len(uuid) == 0 {
		genUUID, err := random.GenUUIDV4()
		if err != nil {
			logrus.WithField("err", err).Panic("middleware logger gen uuid failure")
		}
		uuid = genUUID
	}

	Record = &loggerField{
		FieldLogger: log.WithFields(logrus.Fields{
			// add log fields if you need add more.
			"RequestUUID": uuid,
		}),
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
