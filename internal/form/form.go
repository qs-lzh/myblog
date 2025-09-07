package form

import ()

type FormInterface interface {
	GetValidator() *Validator
}

type CreateForm struct {
	Title   string
	Content string
	Validator
}

func NewCreateForm() *CreateForm {
	return &CreateForm{
		Validator: *NewValidator(),
	}
}

func (form *CreateForm) GetValidator() *Validator {
	return &form.Validator
}
