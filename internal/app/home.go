package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.Logger.LogRequest(r)

	page := "home"
	tmpldata := app.NewTemplateData(r)
	app.render(w, page, tmpldata)
}
