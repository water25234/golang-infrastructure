package shortener

// Business describe linebot business service function
type Business interface {

	// GetShortenerURL mean
	GetShortenerURL(shortenerID string) (shortenerURL string, err error)
}
