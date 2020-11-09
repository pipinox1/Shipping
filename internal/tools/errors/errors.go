package errors

import (
	"fmt"
)

// - Algunos errors comunes en el sistema -

// Unauthorized el usuario no esta autorizado al recurso
var Unauthorized = NewCustom(401, "Unauthorized")

// NotFound cuando un registro no se encuentra en la db
var NotFound = NewCustom(404, "Document not found")

// Internal esta aplicación no sabe como manejar el error
var Internal = NewCustom(500, "Internal server error")

// NewCustom creates a new errCustom
func NewCustom(status int, message string) Custom {
	return &errCustom{
		status:  status,
		Message: message,
	}
}

// NewCustom creates a new errCustom
func NewBusinessError(message string) Custom {
	return &errCustom{
		status:  400,
		Message: message,
	}
}

// Custom es una interfaz para definir errores custom
type Custom interface {
	Status() int
	Error() string
}

// errCustom es un error personalizado para http
type errCustom struct {
	status  int
	Message string `json:"error"`
}

func (e *errCustom) Error() string {
	return fmt.Sprintf(e.Message)
}

// Status http status code
func (e *errCustom) Status() int {
	return e.status
}

//Rest Layer error
type RestClientError struct {
	Message    string
	StatusCode int
}

func (r *RestClientError) Error() string {
	return r.Message
}

func NewRestError(message string, statusCode int) error {
	return &RestClientError{Message: message, StatusCode: statusCode}
}
