package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.Logger.LogRequest(r)

	todos, err := app.TodoModel.GetAll()
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "failed get all todos from database")
	}

	tmpldata := app.NewTemplateData(r)
	tmpldata.Todos = todos

	app.render(w, "home", tmpldata)
}
