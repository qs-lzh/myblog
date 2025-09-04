package form

import "strings"

type Validator struct {
	FieldErrors    map[string]string
	NonFieldErrors []string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

type ValidatorFunc func(string) bool

func (v *Validator) CheckField(fieldContent string, fieldName string, f ValidatorFunc, errMessage string) {
	if !f(fieldContent) {
		v.AddFieldError(fieldName, errMessage)
	}
}

func (v *Validator) AddFieldError(fieldName string, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[fieldName]; !exists {
		v.FieldErrors[fieldName] = message
	}

}

func (v *Validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *Validator) NotBlank(s string) bool {
	return !(len(strings.TrimSpace(s)) == 0)
}
