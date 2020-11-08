package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respondwithError return error message
func ErrorResponse(w http.ResponseWriter, code int, msg string) {
	WebResponse(w, code, map[string]string{"message": msg})
}

// WebResponse write json response format
func WebResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
