package shortener

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	coreRedis "github.com/water25234/golang-infrastructure/core/storage/redis"
	"github.com/water25234/golang-infrastructure/repository/shortener"
)

// New mean shortener.Business by interface
func New(db *sqlx.DB, redis coreRedis.Service) Business {
	return &imple{
		redis:         redis,
		shortenerRepo: shortener.New(db),
	}
}

type imple struct {
	redis         coreRedis.Service
	shortenerRepo shortener.ShortenerRepo
}

var (
	// ErrshortenerIDIsEmpty describe shortener id is empty
	ErrshortenerIDIsEmpty = fmt.Errorf("shortener id is empty")
)

func (im *imple) GetShortenerURL(shortenerID string) (URLEncode string, err error) {
	if len(shortenerID) == 0 {
		return "", ErrshortenerIDIsEmpty
	}

	URLEncode, err = im.shortenerRepo.GetShortenerByID(shortenerID)
	if err != nil {
		return "", err
	}

	decaysecond := 900
	err = im.redis.Set("ShortenerURL:"+shortenerID, URLEncode, decaysecond)
	if err != nil {
		return "", err
	}

	return URLEncode, nil
}
