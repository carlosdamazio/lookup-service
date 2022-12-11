package httputils

import "github.com/kataras/iris/v12"

func RespondError(ctx iris.Context, err error, status int) {
	ctx.StatusCode(status)
	ctx.JSON(iris.Map{
		"message": err.Error(),
	})
}
