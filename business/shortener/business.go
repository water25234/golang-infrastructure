package shortener

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/water25234/golang-infrastructure/repository/shortener"
)

// New mean shortener.Business by interface
func New(DB *sqlx.DB) Business {
	return &imple{
		shortenerRepo: shortener.New(DB),
	}
}

type imple struct {
	shortenerRepo shortener.ShortenerRepo
}

var (
	// ErrshortenerIDIsEmpty describe shortener id is empty
	ErrshortenerIDIsEmpty = fmt.Errorf("shortener id is empty")
)

func (im *imple) GetShortenerURL(shortenerID string) (shortenerURL string, err error) {
	if len(shortenerID) == 0 {
		return "", ErrshortenerIDIsEmpty
	}

	getShortenerID, err := im.shortenerRepo.GetShortenerByID(shortenerID)
	if err != nil {
		return "", err
	}
	return getShortenerID, nil
}
