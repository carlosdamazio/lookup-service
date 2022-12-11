package lookup

import (
	"net"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"github.com/carlosdamazio/lookup-service/internal/models"
)

type Service struct {
	db       *gorm.DB
	lookupFn func(host string) ([]net.IP, error)
}

func New() *Service {
	return &Service{}
}

func (s *Service) WithTx(db *gorm.DB) *Service {
	s2 := s
	s2.db = db
	return s2
}

func (s *Service) WithLookupFn(fn func(host string) ([]net.IP, error)) *Service {
	s2 := s
	s2.lookupFn = fn
	return s2
}

func (s *Service) List() ([]*models.Query, error) {
	var res []*models.Query
	if err := s.db.Order("created_at DESC").Limit(20).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) Lookup(domain, clientIP string) (*models.Query, error) {
	var ips []string

	allIps, err := s.lookupFn(domain)
	if err != nil {
		return nil, err
	}

	for _, ip := range allIps {
		if ipv4 := ip.To4(); ipv4 != nil {
			ips = append(ips, ipv4.String())
		}
	}

	res := &models.Query{
		ID:        uuid.Must(uuid.NewV4()),
		Domain:    domain,
		ClientIP:  clientIP,
		Addresses: ips,
		CreatedAt: time.Now(),
	}

	if err := s.db.Create(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
