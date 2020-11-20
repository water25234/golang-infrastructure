package postgresql

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

// Run mean
func (db *DBPostgresql) Run() (err error) {
	registerNew := postgresql.New()

	err = registerNew.Init()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("register postgresql DB execute run is failure")
		return err
	}

	db.DB = registerNew.GetStorage().(*sqlx.DB)
	return nil
}

// Get means
func (db *DBPostgresql) Get() interface{} {
	return db.DB
}
