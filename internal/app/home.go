package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Data struct {
	Author string
}

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page := "home"
	data := &Data{
		Author: "qs-lzh",
	}
	err := app.render(w, r, page, data)
}
