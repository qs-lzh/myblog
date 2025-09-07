package util

import (
	"time"
)

// parse string s (format "2006-01-02") into time.Time
func ParseDate(s string) time.Time {
	// html输入的日期非空且一定能被正确解析，因此这里不考虑error
	layout := "2006-01-02"
	dueDate, _ := time.Parse(layout, s)

	return dueDate
}
