package app

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) ShowAbout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles(
		"ui/html/layout.tmpl",
		"ui/html/content.tmpl",
		"ui/html/footer.tmpl",
		"ui/html/header.tmpl",
		"ui/html/home.tmpl",
		"ui/html/layout.tmpl",
		"ui/html/title.tmpl",
	)
	if err != nil {
		app.ErrLog.Fatal(err)
	}

	err = tmpl.Execute(w, "")

}
