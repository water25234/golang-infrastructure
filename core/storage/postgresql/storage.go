package postgresql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/water25234/golang-infrastructure/config"
	"github.com/water25234/golang-infrastructure/core/storage"
)

var (
	// ErrNilDBServDB describe DBserv DB struct is nil
	ErrNilDBServDB = fmt.Errorf("DBserv DB struct is nil")
)

var (
	// Storage mean
	Storage DBServ
)

// DBServ mean
type DBServ struct {
	DB *sqlx.DB
}

// New mean
func New() storage.Storage {
	return &DBServ{}
}

// Init mean
func (dbServ *DBServ) Init() (err error) {
	uri := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetAppConfig().DBHost,
		config.GetAppConfig().DBPort,
		config.GetAppConfig().DBUsername,
		config.GetAppConfig().DBPassword,
		config.GetAppConfig().DBDatabase,
	)

	dbServ.DB, err = sqlx.Open(config.GetAppConfig().DBConnection, uri)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"uri":  uri,
			"time": time.Now().UnixNano() / time.Millisecond.Nanoseconds(),
			"err":  err,
		}).Panic("connect postgresql error")
		return err
	}

	err = dbServ.DB.Ping()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"uri": uri,
			"err": err,
		}).Error("failure to connect database")
		return err
	}

	return nil
}

// GetStorage mean
func (dbServ *DBServ) GetStorage() interface{} {
	return dbServ.DB
}

// Disconnect mean
func (dbServ *DBServ) Disconnect() {
	dbServ.DB.Close()
}
