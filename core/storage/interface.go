package storage

// Storage mean
type Storage interface {
	// Init means
	Init() (err error)

	// GetStorage means
	GetStorage() interface{}

	// Disconnect means
	Disconnect()
}
