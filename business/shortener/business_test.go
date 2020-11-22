package shortener

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mRedis "github.com/water25234/golang-infrastructure/core/storage/redis/mocks"
	mRepo "github.com/water25234/golang-infrastructure/repository/shortener/mocks"
)

type ubSuite struct {
	suite.Suite
	shortenerRepo     *mRepo.ShortenerRepo
	redis             *mRedis.Service
	shortenerBusiness Business
}

func TestShortenerBusines(t *testing.T) {
	suite.Run(t, new(ubSuite))
}

func (t *ubSuite) SetupSuite() {
}

func (t *ubSuite) SetupTest() {
	t.shortenerRepo = &mRepo.ShortenerRepo{}
	t.redis = &mRedis.Service{}
	t.shortenerBusiness = &imple{
		redis:         t.redis,
		shortenerRepo: t.shortenerRepo,
	}
}

func (t *ubSuite) TearDownTest() {
	t.shortenerRepo.AssertExpectations(t.T())
}

func (t *ubSuite) TestGetShortenerURL() {
	testCase := []struct {
		Desc        string
		ShortenerID string
		TestFunc    func()
		ExpRes      string
		ExpErr      error
	}{
		{
			Desc:        "TestGetShortenerURL: get shortener URL is success",
			ShortenerID: "lVqELK",
			TestFunc: func() {
				t.shortenerRepo.On("GetShortenerByID", "lVqELK").Return("https://google.com/lVqELK", nil).Once()
				t.redis.On("Set", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil).Once()
			},
			ExpRes: "https://google.com/lVqELK",
			ExpErr: nil,
		},
		{
			Desc:        "TestGetShortenerURL: ShortenerID parameter is emtpy",
			ShortenerID: "",
			TestFunc:    func() {},
			ExpRes:      "",
			ExpErr:      ErrshortenerIDIsEmpty,
		},
		{
			Desc:        "TestGetShortenerURL: get shortener URL from DB is failure",
			ShortenerID: "VqEL",
			TestFunc: func() {
				t.shortenerRepo.On("GetShortenerByID", "VqEL").Return("", fmt.Errorf("sql: no rows in result set"))
			},
			ExpRes: "",
			ExpErr: fmt.Errorf("sql: no rows in result set"),
		},
		{
			Desc:        "TestGetShortenerURL: set redis is failure",
			ShortenerID: "lVqELK",
			TestFunc: func() {
				t.shortenerRepo.On("GetShortenerByID", "lVqELK").Return("https://google.com/lVqELK", nil).Once()
				t.redis.On("Set", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(fmt.Errorf("redis: no rows in result set")).Once()
			},
			ExpRes: "",
			ExpErr: fmt.Errorf("redis: no rows in result set"),
		},
	}

	for _, c := range testCase {
		c.TestFunc()
		result, err := t.shortenerBusiness.GetShortenerURL(c.ShortenerID)
		t.Equal(c.ExpRes, result, c.Desc)
		t.Equal(c.ExpErr, err, c.Desc)
		t.TearDownTest()
	}
}
