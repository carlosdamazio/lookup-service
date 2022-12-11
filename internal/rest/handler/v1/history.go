package v1

import (
	"github.com/kataras/iris/v12"

	"github.com/carlosdamazio/lookup-service/internal/httputils"
)

func (h *LookUpHandler) History(ctx iris.Context) {
	res, err := h.lookupSvc.List()
	if err != nil {
		httputils.RespondError(ctx, err, iris.StatusInternalServerError)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(res)
}
