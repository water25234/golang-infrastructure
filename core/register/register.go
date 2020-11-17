package register

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	RegisterServMap = make(map[string]*RgtrConfig)

	// ErrNilRegisterService describe register service is exists
	ErrNilRegisterService = fmt.Errorf("register parameter is nil")

	// ErrHasBeenRegisted describe register service is exists
	ErrHasBeenRegisted = fmt.Errorf("register service has been registed")

	// ErrNotRegister describe it' not register service
	ErrNotRegister = fmt.Errorf("it' not register service")
)

// Register mean
func Register(registerServName string, rs RgtrService) (err error) {
	if len(registerServName) == 0 || rs == nil {
		panic(ErrNilRegisterService)
	}

	if _, ok := RegisterServMap[registerServName]; ok {
		logrus.WithFields(logrus.Fields{
			"registerServName": registerServName,
		}).Error(ErrHasBeenRegisted)
		return nil
	}

	RegisterServMap[registerServName] = &RgtrConfig{
		Register: rs,
		Once:     &sync.Once{},
	}

	return nil
}

// Run means
func Run(registerServName string) (err error) {
	if _, ok := RegisterServMap[registerServName]; !ok {
		return ErrNotRegister
	}

	RegisterServMap[registerServName].Once.Do(func() {
		RegisterServMap[registerServName].Register.Run()
	})

	return nil
}

// Get means
func Get(registerServName string) (interface{}, error) {
	if _, ok := RegisterServMap[registerServName]; !ok {
		return nil, ErrNotRegister
	}

	return RegisterServMap[registerServName].Register.Get(), nil
}
