package app

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
)

func setLogger(app *iris.Application) {
	app.Logger().SetOutput(os.Stdout)
	app.Logger().SetLevel(os.Getenv("APP_LOG_LEVEL"))
}

func getAccessLog(app *iris.Application) *accesslog.AccessLog {
	ac := accesslog.New(os.Stdout)
	ac.SetFormatter(&accesslog.JSON{})
	return ac
}

func GetApp() *iris.Application {
	app := iris.New()
	setLogger(app)
	ac := getAccessLog(app)
	app.UseRouter(ac.Handler)
	return app
}
