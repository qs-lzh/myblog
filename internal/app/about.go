package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) ShowAbout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.Logger.LogRequest(r)

	page := "about"
	data := &Data{
		Author: "liuzihao",
	}
	app.render(w, r, page, data)
}
