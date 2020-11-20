package register

import (
	"sync"
)

// RgtrConfig means
type RgtrConfig struct {
	Register RgtrService
	Once     *sync.Once
}

// RgtrService means
type RgtrService interface {
	// Run means
	Run() (err error)

	// Get means
	Get() interface{}
}
