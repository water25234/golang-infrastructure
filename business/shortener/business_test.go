package shortener

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ubSuite struct {
	suite.Suite
}

func TestShortenerBusines(t *testing.T) {
	suite.Run(t, new(ubSuite))
}

func (t *ubSuite) SetupSuite() {
}

func (t *ubSuite) SetupTest() {
}

func (t *ubSuite) TearDownTest() {
}

func (t *ubSuite) TestGetShortenerURL() {
}
