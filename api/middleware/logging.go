package middleware

import (
	"log"
	"net/http"
)

// Logging that holds stuff
type Logging struct {
	Logger *log.Logger
}

// Middleware logs info about requests
func (m *Logging) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		m.Logger.Println(req.RequestURI)
		next.ServeHTTP(res, req)
	})
}
