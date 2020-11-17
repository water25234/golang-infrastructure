package register

import (
	"sync"
)

// RgtrConfig mean
type RgtrConfig struct {
	Register RgtrService
	Once     *sync.Once
}

// RgtrService mean
type RgtrService interface {
	Run() (err error)
	Get() interface{}
}
