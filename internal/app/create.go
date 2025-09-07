package app

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/form"
)

func (app *Application) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 后面考虑把所有handler开头和结尾的log写到middleware中
	app.Logger.LogRequest(r)

	data := app.NewTemplateData(r)
	data.Form = form.NewCreateForm()
	app.render(w, "create", data)

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

	createForm.CheckField(createForm.NotBlank(createForm.Title), "title", "title field should not be blank!")
	createForm.CheckField(createForm.MaxLength(createForm.Title, 255), "title", "title field shoule be no more than 255 runes")
	createForm.CheckField(createForm.NotBlank(createForm.Content), "content", "content field should not be blank!")

	data := app.NewTemplateData(r)
	data.Form = createForm

	if !createForm.Valid() {
		app.render(w, "create", data)
		return
	}

	err = app.TodoModel.Insert(createForm.Title, createForm.Content, time.Now())
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "failed to insert into database")
	}
	app.Logger.LogDBModify("Insert", "todos")

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
