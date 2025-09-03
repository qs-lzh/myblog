package app

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Application struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, page string, data any) {
	templateSet, err := template.ParseFiles("./ui/html/base.tmpl")
	if err != nil {
		app.ErrLog.Fatal(err)
	}

	filePath := filepath.Join("./ui/html/pages", page+".tmpl")

	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
		app.ErrLog.Fatal(err)
	}

	templateSet, err = templateSet.ParseFiles(filePath)
	if err != nil {
		app.ErrLog.Fatal(err)
	}

	err = templateSet.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.ErrLog.Fatal(err)
	}
}
