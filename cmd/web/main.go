package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/app"
	"github.com/qs-lzh/myblog/internal/errors"
	"github.com/qs-lzh/myblog/internal/logger"
)

func main() {
	logger := logger.NewLogger()
	errorHandler := errors.NewErrorHandler(logger)
	sessionManager := scs.New()
	sessionManager.Lifetime = time.Hour * 5

	app := &app.Application{
		Logger:         logger,
		ErrorHandler:   errorHandler,
		SessionManager: sessionManager,
	}

	router := httprouter.New()

	router.GET("/home", app.ShowHome)
	router.GET("/create", app.Create)
	router.POST("/create", app.CreatePost)

	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.SessionManager.LoadAndSave(router),
	}

	err := srv.ListenAndServe()
	if err != nil {
		app.ErrorHandler.Logger.ErrLog.Fatal(err)
	}
}
