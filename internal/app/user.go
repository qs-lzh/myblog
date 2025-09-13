package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/qs-lzh/myblog/internal/form"
)

func (app *Application) UserSignup(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	app.Logger.LogRequest(r)

	data := app.NewTemplateData(r)
	app.render(w, "signup", data)

	return
}

func (app *Application) UserSignupPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	app.Logger.LogRequest(r)

	err := r.ParseForm()
	if err != nil {
		app.ErrorHandler.UnProcessableEntity(w, "")
	}

	signupForm := form.NewSignupForm()
	signupForm.Email = r.FormValue("email")
	signupForm.Password = r.FormValue("password")
	signupForm.ConfirmPassword = r.FormValue("confirm_password")

	signupForm.CheckField(signupForm.IsEmail(signupForm.Email), "email", "the email address is not valid")
	signupForm.CheckField(signupForm.IsSame(signupForm.ConfirmPassword, signupForm.Password), "confirm_password", "confirm_password should be the same as password")
	signupForm.CheckField(signupForm.MinLength(signupForm.Password, 6), "password", "the password should be no less than 6 characters")

	data := app.NewTemplateData(r)
	data.Form = signupForm

	if !signupForm.Valid() {
		app.render(w, "signup", data)
		return
	}
}

func (app *Application) UserLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	app.Logger.LogRequest(r)

	data := app.NewTemplateData(r)
	app.render(w, "login", data)

	return
}

func (app *Application) UserLoginPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	app.Logger.LogRequest(r)

	err := r.ParseForm()
	if err != nil {
		app.ErrorHandler.UnProcessableEntity(w, "")
	}

	loginForm := form.NewLoginForm()
	loginForm.Email = r.FormValue("email")
	loginForm.Password = r.FormValue("password")

	loginForm.CheckField(loginForm.IsEmail(loginForm.Email), "email", "the email address is not valid")
	loginForm.CheckField(loginForm.MinLength(loginForm.Password, 6), "password", "the password should be no less than 6 characters")

	data := app.NewTemplateData(r)
	data.Form = loginForm

	if !loginForm.Valid() {
		app.render(w, "login", data)
		return
	}
}

func (app *Application) ShowUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	app.Logger.LogRequest(r)

	data := app.NewTemplateData(r)

	app.render(w, "user", data)
}
