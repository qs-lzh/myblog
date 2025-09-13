package form

import (
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"
)

type validatorInterface interface {
	Valid() bool
	CheckField(fieldContent string, fieldName string, f validatorFunc, errMessage string)
	AddFieldError(fieldName string, message string)
	AddNonFieldError()
}

type validator struct {
	FieldErrors    map[string]string
	NonFieldErrors []string
}

func newValidator() *validator {
	return &validator{
		FieldErrors:    make(map[string]string),
		NonFieldErrors: make([]string, 0),
	}
}

func (v *validator) Valid() bool {
	return len(v.FieldErrors) == 0 && len(v.NonFieldErrors) == 0
}

type validatorFunc func(string) bool

// if ok is false, then Validator.FieldErrors[fieldName] = errMessage
func (v *validator) CheckField(ok bool, fieldName string, errMessage string) {
	if !ok {
		v.AddFieldError(fieldName, errMessage)
	}
}

func (v *validator) AddFieldError(fieldName string, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[fieldName]; !exists {
		v.FieldErrors[fieldName] = message
	}

}

func (v *validator) AddNonFieldError(message string) {
	v.NonFieldErrors = append(v.NonFieldErrors, message)
}

func (v *validator) NotBlank(s string) bool {
	return !(len(strings.TrimSpace(s)) == 0)
}

func (v *validator) MaxLength(s string, maxLength int) bool {
	return utf8.RuneCountInString(s) <= maxLength
}

func (v *validator) MinLength(s string, minLength int) bool {
	return utf8.RuneCountInString(s) >= minLength
}

func (v *validator) AfterNow(t time.Time) bool {
	return t.After(time.Now())
}

func (v *validator) IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (v *validator) IsSame(s1 string, s2 string) bool {
	return s1 == s2
}
