package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/form"
)

func (app *Application) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 后面考虑把所有handler开头和结尾的log写到middleware中
	app.Logger.LogRequest(r)

	data := &Data{}
	app.render(w, r, "create", data)

	app.Logger.LogRequest(r)
}

func (app *Application) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.Logger.LogRequest(r)

	err := r.ParseForm()
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "failed parse form")
		return
	}

	createForm := &form.CreateForm{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	createForm.CheckField(createForm.Title, "title", createForm.NotBlank, "title field should not be blank!")

	if !createForm.Valid() {
		app.SessionManager.Put(r.Context(), "flash", "form input wront!")
		data := app.NewTemplateData(r)
		app.render(w, r, "create", data)
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)

	app.Logger.LogRequest(r)
}
