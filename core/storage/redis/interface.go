package redis

// Service meas
type Service interface {

	// Init means
	Init() (err error)

	// GetStorage means
	GetStorage() interface{}

	// Disconnect means
	Disconnect()

	// Set
	Set(key string, value string, num int) (err error)

	// Get
	Get(key string) (val string, err error)

	// Del
	Del(key string) (err error)

	// Keys
	Keys(key string) (val []string, err error)
}
