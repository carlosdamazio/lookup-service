package binder

import (
	"github.com/kataras/iris/v12"

	v1 "github.com/carlosdamazio/lookup-service/internal/rest/handler/v1"
)

func BindAll(application *iris.Application) {
	v1Handler := v1.NewLookupHandler().WithSvc()
	v1Handler.BindV1(application)
}
