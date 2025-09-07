package data

import (
	"time"

	"github.com/qs-lzh/myblog/internal/form"
)

type TemplateData struct {
	Author string
	Time   time.Time
	Flash  string
	Form   form.FormInterface
}
