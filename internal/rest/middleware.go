package rest

import (
	"net/http"
	"os"
	"strings"
)

// User es el usuario logueado
type User struct {
	ID          string   `json:"id"  validate:"required"`
	Name        string   `json:"name"  validate:"required"`
	Permissions []string `json:"permissions"`
	Login       string   `json:"login"  validate:"required"`
}

func securityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if strings.Index(tokenString, "bearer ") != 0 {
			ErrorResponse(w, 401, "Unauthorized")
			return
		}
		token := tokenString[7:]

		if os.Getenv("SKIP_AUTH") != "true"{
			req, err := http.NewRequest("GET",os.Getenv("USER_URL")+"/v1/users/current", nil)
			if err != nil || req == nil{
				ErrorResponse(w, 500, "internal_server_error")
				return
			}
			req.Header.Add("Authorization", "bearer "+token)
			resp, err := http.DefaultClient.Do(req)
			if err != nil || resp.StatusCode != 200 {
				ErrorResponse(w, 401, "Unauthorized")
				return
			}
		}
	
		next.ServeHTTP(w, r)
})
}
