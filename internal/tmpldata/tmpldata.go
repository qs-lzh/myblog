package data

import (
	"github.com/qs-lzh/myblog/internal/data"
	"github.com/qs-lzh/myblog/internal/form"
)

type TemplateData struct {
	Author string
	Flash  string
	Form   form.FormInterface
	Todo   *data.Todo
	Todos  []*data.Todo
}
