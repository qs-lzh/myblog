package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) DeletePost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	app.Logger.LogRequest(r)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.ErrorHandler.ServerError(w, err, "failed to convert id to type string")
		return
	}

	err = app.TodoModel.Delete(id)
	if err != nil {
		app.ErrorHandler.ServerError(w, err, fmt.Sprintf("failed to delete todo id = %d", id))
		return
	}

	app.SessionManager.Put(r.Context(), "flash", fmt.Sprintf("delete todo id = %d successfully", id))
	app.Logger.LogDBModify("delete", "todos")

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
