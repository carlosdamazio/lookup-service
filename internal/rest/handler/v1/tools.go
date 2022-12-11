package v1

import (
	"errors"

	"github.com/carlosdamazio/lookup-service/internal/httputils"
	"github.com/kataras/iris/v12"
)

func (h *LookUpHandler) Lookup(ctx iris.Context) {
	domain := ctx.URLParam("domain")
	if domain == "" {
		err := errors.New("domain URL param is required")
		httputils.RespondError(ctx, err, iris.StatusBadRequest)
		return
	}

	clientIP := ctx.Request().RemoteAddr

	res, err := h.lookupSvc.Lookup(domain, clientIP)
	if err != nil {
		httputils.RespondError(ctx, err, iris.StatusInternalServerError)
		return
	}

	if len(res.Addresses) == 0 {
		httputils.RespondError(ctx, errors.New("resource not found"), iris.StatusNotFound)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(res)
}

func (h *LookUpHandler) Validate(ctx iris.Context) {
	var req ValidateIPRequest
	if err := ctx.ReadJSON(&req); err != nil {
		httputils.RespondError(ctx, err, iris.StatusBadRequest)
		return
	} else if req.IP == "" {
		err := errors.New("IP is required")
		httputils.RespondError(ctx, err, iris.StatusBadRequest)
		return
	}

	ctx.StatusCode(iris.StatusOK)

	if err := req.Validate(); err != nil {
		ctx.JSON(iris.Map{
			"status": false,
		})
		return
	}

	ctx.JSON(iris.Map{
		"status": true,
	})
}
