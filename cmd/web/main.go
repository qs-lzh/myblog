package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/app"
	"github.com/qs-lzh/myblog/internal/errors"
	"github.com/qs-lzh/myblog/internal/logger"
)

func main() {
	logger := logger.NewLogger()
	errorHandler := errors.NewErrorHandler(logger)

	app := &app.Application{
		Logger:       logger,
		ErrorHandler: errorHandler,
	}

	router := httprouter.New()

	router.GET("/home", app.ShowHome)

	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		app.ErrorHandler.Logger.ErrLog.Fatal()
	}
}
