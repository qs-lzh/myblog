package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/app"
	"github.com/qs-lzh/myblog/internal/errors"
)

func main() {

	errorHandler := errors.NewErrorHandler()

	app := &app.Application{
		ErrorHandler: *errorHandler,
	}

	router := httprouter.New()

	router.GET("/home", app.ShowHome)
	router.GET("/about", app.ShowAbout)

	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		app.ErrorHandler.ErrLog.Fatal(err)
	}
}
