package v1

import (
	"fmt"
	"net"
	"regexp"

	"github.com/kataras/iris/v12"

	"github.com/carlosdamazio/lookup-service/internal/db"
	"github.com/carlosdamazio/lookup-service/internal/models"
	"github.com/carlosdamazio/lookup-service/internal/service/lookup"
)

type (
	lookupService interface {
		List() ([]*models.Query, error)
		Lookup(domain, clientIP string) (*models.Query, error)
	}
	LookUpHandler struct {
		lookupSvc lookupService
	}
)

func NewLookupHandler() *LookUpHandler {
	return &LookUpHandler{}
}

func (h *LookUpHandler) BindV1(app *iris.Application) {
	party := app.Party("/v1")
	party.Get("/history", h.History)
	party.Get("/tools/lookup", h.Lookup)
	party.Post("/tools/validate", h.Validate)
}

func (h *LookUpHandler) WithSvc() *LookUpHandler {
	h2 := h
	h2.lookupSvc = lookup.New().WithTx(db.GetDB()).WithLookupFn(net.LookupIP)
	return h2
}

type ValidateIPRequest struct {
	IP string `json:"ip" validate:"required"`
}

func (r *ValidateIPRequest) Validate() error {
	if ok, err := regexp.MatchString(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`, r.IP); !ok || err != nil {
		return fmt.Errorf("IP \"%s\" is not a valid address", r.IP)
	}
	return nil
}

type ValidateIPResponse struct {
	Status bool `json:"status"`
}
