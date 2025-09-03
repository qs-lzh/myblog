package app

import (
	"log"
)

type Application struct {
	InfoLog *log.Logger
	ErrLog  *log.Logger
}
