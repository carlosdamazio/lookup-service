package v1

import (
	"errors"
	"testing"

	"github.com/carlosdamazio/lookup-service/internal/models"
	"github.com/carlosdamazio/lookup-service/internal/rest/handler/server/app"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/mock"
)

func TestValidateNoBody(t *testing.T) {
	app := app.GetApp()

	handler := &LookUpHandler{}
	handler.BindV1(app)

	e := httptest.New(t, app)

	e.POST("/v1/tools/validate").WithJSON(iris.Map{}).Expect().Status(httptest.StatusBadRequest)
}

func TestValidateInvalidIPV4Addr(t *testing.T) {
	app := app.GetApp()

	handler := &LookUpHandler{}
	handler.BindV1(app)

	e := httptest.New(t, app)
	e.POST("/v1/tools/validate").WithJSON(ValidateIPRequest{IP: "10.256.0.100020"}).Expect().Status(httptest.StatusOK)
}

func TestValidateOk(t *testing.T) {
	app := app.GetApp()

	handler := &LookUpHandler{}
	handler.BindV1(app)

	e := httptest.New(t, app)
	e.POST("/v1/tools/validate").WithJSON(ValidateIPRequest{IP: "10.0.0.1"}).Expect().Status(httptest.StatusOK)
}

func TestLookUpNoDomain(t *testing.T) {
	app := app.GetApp()

	handler := &LookUpHandler{}
	handler.BindV1(app)

	e := httptest.New(t, app)
	e.GET("/v1/tools/lookup").Expect().Status(httptest.StatusBadRequest)
}

func TestLookUpError(t *testing.T) {
	app := app.GetApp()

	lookupSvc := newLookupService(t)
	lookupSvc.On("Lookup", mock.Anything, mock.Anything).Return(nil, errors.New("hey"))

	handler := &LookUpHandler{lookupSvc: lookupSvc}
	handler.BindV1(app)

	e := httptest.New(t, app)
	e.GET("/v1/tools/lookup").
		WithQuery("domain", "damazio.dev").
		Expect().
		Status(httptest.StatusInternalServerError)
}

func TestLookUpOk(t *testing.T) {
	app := app.GetApp()

	lookupSvc := newLookupService(t)
	lookupSvc.On("Lookup", mock.Anything, mock.Anything).Return(&models.Query{
		Addresses: []string{"127.0.0.1"},
	}, nil)

	handler := &LookUpHandler{lookupSvc: lookupSvc}
	handler.BindV1(app)

	e := httptest.New(t, app)
	e.GET("/v1/tools/lookup").
		WithQuery("domain", "damazio.dev").
		Expect().
		Status(httptest.StatusOK)
}

func TestLookUpNotFound(t *testing.T) {
	app := app.GetApp()

	lookupSvc := newLookupService(t)
	lookupSvc.On("Lookup", mock.Anything, mock.Anything).Return(&models.Query{
		Addresses: []string{},
	}, nil)

	handler := &LookUpHandler{lookupSvc: lookupSvc}
	handler.BindV1(app)

	e := httptest.New(t, app)
	e.GET("/v1/tools/lookup").
		WithQuery("domain", "damazio.dev").
		Expect().
		Status(httptest.StatusNotFound)
}
