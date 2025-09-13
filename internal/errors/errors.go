package errors

import (
	"fmt"
	"net/http"

	"github.com/qs-lzh/myblog/internal/logger"
)

type ErrorHandler struct {
	Logger *logger.Logger
}

func NewErrorHandler(l *logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		Logger: l,
	}
}

func (h *ErrorHandler) ServerError(w http.ResponseWriter, err error, msg string) {
	h.Logger.Error(fmt.Sprintf("Server Error - %s: %v", msg, err))
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (h *ErrorHandler) ClientError(w http.ResponseWriter, statusCode int, msg string) {
	h.Logger.Error(fmt.Sprintf("Client Error - (%d): %s", statusCode, msg))
	http.Error(w, "Client Error", statusCode)
}

/*
  You should NEVER use ClientError() UNLESS the specialized errors doesn't satisfy your need
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

func (h *ErrorHandler) UnProcessableEntity(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Unprocessable Entity"
	}
	h.ClientError(w, http.StatusUnprocessableEntity, msg)
}
