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

func (app *Application) render(w http.ResponseWriter, page string, data any) {
	// 解析base.tmpl和partials/中的模版
	templateSet, err := template.ParseFiles("./ui/html/base.tmpl")
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "Not found ./ui/html/base.tmpl")
		return
	}
	dir := "./ui/html/partials/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		app.ErrorHandler.ServerError(w, err, fmt.Sprintf("Not found director %s", dir))
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := dir + entry.Name()
			fmt.Println(filePath)
			templateSet, err = templateSet.ParseFiles(filePath)
			if err != nil {
				app.ErrorHandler.ServerError(w, err, fmt.Sprintf("failed to parse %s", filePath))
			}
		}
	}

	// 解析pages/中的页面模版
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

	// 执行模版
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
