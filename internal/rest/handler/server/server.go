package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/carlosdamazio/lookup-service/internal/db"
	"github.com/carlosdamazio/lookup-service/internal/rest/handler/server/app"
	"github.com/carlosdamazio/lookup-service/internal/rest/handler/server/binder"
)

func StartServer(addr, port string) {
	app := app.GetApp()
	binder.BindAll(app)

	app.Get("/", func(c iris.Context) {
		_ = c.JSON(iris.Map{
			"version":    os.Getenv("APP_VERSION"),
			"date":       time.Now().Unix(),
			"kubernetes": "kubernetes" == os.Getenv("APP_MANAGED_BY"),
		})
	})

	app.Get("/health", func(c iris.Context) {
		_ = c.JSON(iris.Map{
			"status": "OK",
		})
	})

	// prometheus metrics
	app.Get("/metrics", iris.FromStd(promhttp.Handler()))

	idleConns := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		app.Logger().Info("server shutting down graciously")

		timeout := 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := app.Shutdown(ctx); err != nil {
			app.Logger().Fatal("http server shutdown", "error", err)
		}

		if o, err := db.GetDB().DB(); err != nil {
			app.Logger().Fatal("get DB for shutdown", "error", err)
		} else {
			if err := o.Close(); err != nil {
				app.Logger().Fatal("closing DB connection object", "error", err)
			}
		}

		time.Sleep(timeout)
		close(idleConns)
	})

	if err := app.Listen(fmt.Sprintf("%s:%s", addr, port), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		app.Logger().Fatal("http server listening", "error", err)
	}

	<-idleConns
}
