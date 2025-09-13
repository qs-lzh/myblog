package tmpldata

import (
	"github.com/qs-lzh/myblog/internal/data"
	"github.com/qs-lzh/myblog/internal/form"
)

type TemplateData struct {
	AA    int
	Flash string
	Form  form.FormInterface
	Todo  *data.Todo
	Todos []*data.Todo
	User  data.User
}
