package register

// RgtrType mean
type RgtrType int32

const (
	// RegisterPostgresql means
	RegisterPostgresql RgtrType = 1

	// RegisterRedis means
	RegisterRedis RgtrType = 2

	// RegisterMySQL means
	RegisterMySQL RgtrType = 3

	// RegisterMongoDB means
	RegisterMongoDB RgtrType = 4
)

const (
	// StoragePostgresql means
	StoragePostgresql string = "StoragePostgresql"

	// StorageRedis means
	StorageRedis string = "StorageRedis"

	// StorageMySQL means
	StorageMySQL string = "StorageMySQL"

	// StorageMongoDB means
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
