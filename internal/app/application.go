package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/qs-lzh/myblog/internal/data"
	"github.com/qs-lzh/myblog/internal/errors"
	"github.com/qs-lzh/myblog/internal/logger"
)

type Application struct {
	Logger         *logger.Logger
	ErrorHandler   *errors.ErrorHandler
	Data           *Data
	SessionManager *scs.SessionManager
}

type Data struct {
	Author string
	Flash  string
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, page string, data any) {
	templateSet, err := template.ParseFiles("./ui/html/base.tmpl")
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "Not found ./ui/html/base.tmpl")
		return
	}

	filePath := filepath.Join("./ui/html/pages", page+".tmpl")

	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			app.ErrorHandler.ServerError(w, err, fmt.Sprintf("Not found %s", filePath))
			return
		}
		app.ErrorHandler.ServerError(w, err, "os.Stat(filePath) failed")
		return
	}

	templateSet, err = templateSet.ParseFiles(filePath)
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "failed to parse files to")
		return
	}

	err = templateSet.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "failed to execute template")
		return
	}

	app.Logger.LogPageRender(page)
}

func (app *Application) NewTemplateData(r *http.Request) *data.TemplateData {
	return &data.TemplateData{
		Time:  time.Now(),
		Flash: app.SessionManager.PopString(r.Context(), "flash"),
	}
}
