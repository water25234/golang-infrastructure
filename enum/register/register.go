package register

// RgtrType mean
type RgtrType int32

const (
	// RegisterPostgresql mean
	RegisterPostgresql RgtrType = 1

	// RegisterRedis mean
	RegisterRedis RgtrType = 2

	// RegisterMySQL mean
	RegisterMySQL RgtrType = 3

	// RegisterMongoDB mean
	RegisterMongoDB RgtrType = 4
)

const (
	// StoragePostgresql mean
	StoragePostgresql string = "StoragePostgresql"

	// StorageRedis mean
	StorageRedis string = "StorageRedis"

	// StorageMySQL mean
	StorageMySQL string = "StorageMySQL"

	// StorageMongoDB mean
	StorageMongoDB string = "StorageMongoDB"
)

const (
	// InterfaceShortener mean
	InterfaceShortener string = "InterfaceShortener"
)

func (p RgtrType) String() string {
	switch p {
	case RegisterPostgresql:
		return "StoragePostgresql"
	case RegisterRedis:
		return "StorageRedis"
	case RegisterMySQL:
		return "StorageMySQL"
	case RegisterMongoDB:
		return "StorageMongoDB"
	default:
		return "UNKNOWN"
	}
}
