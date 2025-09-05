package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.Logger.LogRequest(r)

	page := "home"
	data := &Data{
		Author: "liuzihao",
	}
	app.render(w, page, data)
}
