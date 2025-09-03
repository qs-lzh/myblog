package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/app"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &app.Application{
		InfoLog: infoLog,
		ErrLog:  errLog,
	}

	router := httprouter.New()

	router.GET("/about", app.ShowAbout)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		app.ErrLog.Fatal(err)
	}
}
