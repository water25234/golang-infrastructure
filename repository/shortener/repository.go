package shortener

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// New mean
func New(db *sqlx.DB) ShortenerRepo {
	return &shortenerRepo{
		db: db,
	}
}

type shortenerRepo struct {
	db *sqlx.DB
}

// shortener mean
type shortener struct {
	ShortenerID int64  `json:"shortener_id" db:"shortener_id"`
	CreateDate  string `json:"create_date" db:"create_date"`
	ModifyDate  string `json:"modify_date" db:"modify_date"`
	URLEncode   string `json:"url_encode" db:"url_encode"`
	URL         string `json:"url" db:"url"`
	IsClose     bool   `json:"is_close" db:"is_close"`
}

func (impl *shortenerRepo) GetShortenerByID(shortenerID string) (URLEncode string, err error) {
	sqlStr := `
		SELECT url_encode, url 
		FROM shortener 
		Where url_encode = :url_encode;
	`

	rows, err := impl.db.NamedQuery(
		sqlStr,
		shortener{
			URLEncode: shortenerID,
		},
	)

	if err != nil {
		logrus.WithField("err", err).Error("query administrator info by ID is failure.")
		return "", err
	}

	shortener := &shortener{}
	if rows.Next() {
		if err := rows.StructScan(&shortener); err != nil {
			logrus.WithField("err", err).Error("query administrator info by ID scan struct is failure.")
			return "", err
		}
	}

	return shortener.URL, nil
}
