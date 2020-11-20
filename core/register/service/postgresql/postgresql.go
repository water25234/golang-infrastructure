package postgresql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/core/register"
	rgtrEnum "github.com/water25234/golang-infrastructure/enum/register"
)

// RegisterDBRun mean
func RegisterDBRun() {
	register.Run(rgtrEnum.StoragePostgresql)
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
