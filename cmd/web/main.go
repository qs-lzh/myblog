package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qs-lzh/myblog/internal/app"
	"github.com/qs-lzh/myblog/internal/data"
	"github.com/qs-lzh/myblog/internal/errors"
	"github.com/qs-lzh/myblog/internal/logger"
)

func main() {
	logger := logger.NewLogger()
	errorHandler := errors.NewErrorHandler(logger)
	sessionManager := scs.New()
	sessionManager.Lifetime = time.Hour * 5

	dsn := flag.String("dsn", "web:mysql123..@unix(/tmp/mysql.sock)/todo?parseTime=true", "MYSQL data source name")

	db, err := data.OpenDB(*dsn)
	if err != nil {
		logger.ErrLog.Fatal(err)
	}
	defer db.Close()

	todoModel := data.TodoModel{DB: db}

	app := &app.Application{
		Logger:         logger,
		ErrorHandler:   errorHandler,
		SessionManager: sessionManager,
		TodoModel:      todoModel,
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

	err = srv.ListenAndServe()
	if err != nil {
		app.ErrorHandler.Logger.ErrLog.Fatal(err)
	}
}
