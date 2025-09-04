package errors

import (
	"log"
	"net/http"
	"os"
)

type ErrorHandler struct {
	infoLog *log.Logger
	ErrLog  *log.Logger
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{
		infoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		ErrLog:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (h *ErrorHandler) ServerError(w http.ResponseWriter, err error, msg string) {
	h.ErrLog.Printf("Server Error - %s: %v", msg, err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (h *ErrorHandler) ClientError(w http.ResponseWriter, statusCode int, msg string) {
	h.ErrLog.Printf("Client Error - (%d): %s", statusCode, msg)
	http.Error(w, "Client Error", statusCode)
}

/*
  You should NEVER use ClientError() UNLESS the specialized errors below doesn't satisfy your need
*/

func (h *ErrorHandler) NotFound(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Page not found"
	}
	h.ClientError(w, http.StatusNotFound, msg)
}

func (h *ErrorHandler) BadRequest(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Invalid request format"
	}
	h.ClientError(w, http.StatusBadRequest, msg)
}

func (h *ErrorHandler) Unauthorized(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "User not authorized"
	}
	h.ClientError(w, http.StatusUnauthorized, msg)
}
