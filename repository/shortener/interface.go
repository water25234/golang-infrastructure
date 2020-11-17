package shortener

// ShortenerRepo mean
type ShortenerRepo interface {
	GetShortenerByID(shortenerID string) (URLEncode string, err error)
}
