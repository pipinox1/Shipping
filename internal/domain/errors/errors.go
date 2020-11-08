package errors

import (
	"encoding/json"
	"fmt"
)

// - Algunos errors comunes en el sistema -

// Unauthorized el usuario no esta autorizado al recurso
var Unauthorized = NewCustom(401, "Unauthorized")

// NotFound cuando un registro no se encuentra en la db
var NotFound = NewCustom(400, "Document not found")

// Internal esta aplicación no sabe como manejar el error
var Internal = NewCustom(500, "Internal server error")

// - Creación de errors -

// NewValidationField crea un error de validación para un solo campo
func NewValidationField(field string, err string) Validation {
	return &errValidation{
		Messages: []errField{
			{
				Path:    field,
				Message: err,
			},
		},
	}
}

// NewValidation crea un error de validación para un solo campo
func NewValidation() Validation {
	return &errValidation{
		Messages: []errField{},
	}
}

// NewCustom creates a new errCustom
func NewCustom(status int, message string) Custom {
	return &errCustom{
		status:  status,
		Message: message,
	}
}

// NewCustom creates a new errCustom
func NewBusinessError( message string) Custom {
	return &errCustom{
		status:  400,
		Message: message,
	}
}


//  - Algunas definiciones necesarias -

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

// Validation es una interfaz para definir errores custom
// Validation es un error de validaciones de parameteros o de campos
type Validation interface {
	Add(path string, message string) Validation
	Size() int
	Error() string
}

// errField define un campo inválido. path y mensaje de error
type errField struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

// ErrValidation es un error de validaciones de parameteros o de campos
type errValidation struct {
	Messages []errField `json:"messages"`
}

func (e *errValidation) Error() string {
	body, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ErrValidation que no se puede pasar a json.")
	}
	return fmt.Sprintf(string(body))
}

// Add agrega errores a un validation error
func (e *errValidation) Add(path string, message string) Validation {
	err := errField{
		Path:    path,
		Message: message,
	}
	e.Messages = append(e.Messages, err)
	return e
}

// Size devuelve la cantidad de errores
func (e *errValidation) Size() int {
	return len(e.Messages)
}
