package form

import (
	"time"
)

type FormInterface interface {
	GetValidator() *validator
}

type CreateForm struct {
	Title   string
	Content string
	DueDate time.Time
	validator
}

func NewCreateForm() *CreateForm {
	return &CreateForm{
		validator: *newValidator(),
	}
}

func (form *CreateForm) GetValidator() *validator {
	return &form.validator
}

type SignupForm struct {
	Email           string
	Password        string
	ConfirmPassword string
	validator
}

func (form *SignupForm) GetValidator() *validator {
	return &form.validator
}

func NewSignupForm() *SignupForm {
	return &SignupForm{
		validator: *newValidator(),
	}
}

type LoginForm struct {
	Email    string
	Password string
	validator
}

func (form *LoginForm) GetValidator() *validator {
	return &form.validator
}

func NewLoginForm() *LoginForm {
	return &LoginForm{
		validator: *newValidator(),
	}
}
