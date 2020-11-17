package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/core/register"
	"github.com/water25234/golang-infrastructure/core/storage/postgresql"
	rgtrEnum "github.com/water25234/golang-infrastructure/enum/register"
)

// DBPostgresql mean
type DBPostgresql struct {
	DB *sqlx.DB
}

func init() {
	register.Register(rgtrEnum.StoragePostgresql, &DBPostgresql{})
}

// RegisterDBRun mean
func RegisterDBRun() {
	register.Run(rgtrEnum.StoragePostgresql)
}

// Run mean
func (db *DBPostgresql) Run() (err error) {
	registerNew := postgresql.New()

	err = registerNew.Init()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("register execute run is failure")
		return err
	}

	db.DB = registerNew.GetStorage()
	return nil
}

// Get means
func (db *DBPostgresql) Get() interface{} {
	return db.DB
}

// GetPostgresqlDB means
func GetPostgresqlDB() *sqlx.DB {
	dbDrive, err := register.Get(rgtrEnum.StoragePostgresql)
	if err != nil {
		logrus.WithField("err", err).Panic("get postgresql DB failure")
		return nil
	}

	dbConn, ok := dbDrive.(*sqlx.DB)
	if !ok {
		logrus.Panic("trasfer db.Drive to sqlx.DB is failure")
		return nil
	}
	return dbConn
}
