package lookup

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/carlosdamazio/lookup-service/internal/db"
	"github.com/carlosdamazio/lookup-service/internal/models"
)

type LookupTestSuite struct {
	suite.Suite
	DB *gorm.DB
}

func TestLookupTestSuite(t *testing.T) {
	suite.Run(t, new(LookupTestSuite))
}

func (t *LookupTestSuite) SetupSuite() {
	t.DB = db.GetDB()
}

func (t *LookupTestSuite) TearDownTest() {
	if err := t.DB.Where("1 = 1").Delete(&models.Query{}).Error; err != nil {
		log.Fatal("teardown test", "error", err)
	}
}

func (t *LookupTestSuite) TestLookupCorrect() {
	lookupSvc := New().WithTx(t.DB).WithLookupFn(func(host string) ([]net.IP, error) {
		return []net.IP{[]byte("192.168.0.1")}, nil
	})

	res, err := lookupSvc.Lookup("damazio.dev", "172.0.1.18")
	t.NoError(err)
	t.NotNil(res)

	var lookups []*models.Query
	t.NoError(t.DB.Find(&lookups).Error)
	t.Len(lookups, 1)
	t.Equal("damazio.dev", lookups[0].Domain)
}
