package service

import (
	"github.com/sirupsen/logrus"

	"github.com/water25234/golang-infrastructure/business/shortener"
	"github.com/water25234/golang-infrastructure/core/register"
	rgtrEnum "github.com/water25234/golang-infrastructure/enum/register"
)

// shortenerInterface mean
type shortenerInterface struct {
	shortenerBiz *shortener.Business
}

func init() {
	register.Register(rgtrEnum.InterfaceShortener, &shortenerInterface{})
}

// RegisterShortenerInterfaceRun mean
func RegisterShortenerInterfaceRun() {
	register.Run(rgtrEnum.InterfaceShortener)
}

// Run mean
func (impl *shortenerInterface) Run() (err error) {

	shortenerBiz := shortener.New(GetPostgresqlDB())

	impl.shortenerBiz = &shortenerBiz

	return nil
}

func (impl *shortenerInterface) Get() interface{} {
	return impl.shortenerBiz
}

// GetShortenerBusiness returns shortener.Business interface
func GetShortenerBusiness() shortener.Business {
	shortenerB, err := register.Get(rgtrEnum.InterfaceShortener)

	if err != nil {
		logrus.WithField("err", err).Panic("get shortener business failure")
		return nil
	}

	shortenerInterface, ok := shortenerB.(*shortener.Business)
	if !ok {
		logrus.Panic("trasfer shortener.Business interface in service layer failure")
		return nil
	}

	return *shortenerInterface
}
