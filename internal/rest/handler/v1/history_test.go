package v1

import (
	"errors"
	"testing"

	"github.com/kataras/iris/v12/httptest"

	"github.com/carlosdamazio/lookup-service/internal/models"
	"github.com/carlosdamazio/lookup-service/internal/rest/handler/server/app"
)

func TestHistoryOk(t *testing.T) {
	app := app.GetApp()

	lookupSvc := newLookupService(t)
	lookupSvc.On("List").Return([]*models.Query{}, nil)

	handler := &LookUpHandler{lookupSvc: lookupSvc}
	handler.BindV1(app)

	e := httptest.New(t, app)

	e.GET("/v1/history").Expect().Status(httptest.StatusOK)
}

func TestHistoryNotOk(t *testing.T) {
	app := app.GetApp()

	lookupSvc := newLookupService(t)
	lookupSvc.On("List").Return(nil, errors.New("hey there"))

	handler := &LookUpHandler{lookupSvc: lookupSvc}
	handler.BindV1(app)

	e := httptest.New(t, app)

	e.GET("/v1/history").Expect().Status(httptest.StatusInternalServerError)
}
