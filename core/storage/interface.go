package storage

import "github.com/jmoiron/sqlx"

// Storage mean
type Storage interface {
	Init() (err error)

	SetStorage() (err error)

	GetStorage() *sqlx.DB
}
