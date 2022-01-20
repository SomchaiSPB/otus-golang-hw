package internalhttp

import (
	"net/http"
	"time"
)

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		s.logger.Sugar().Infof("%s [%s] %s %s %s %s",
			r.RemoteAddr,
			currentTime.Format("02/01/2006:15:04:05 MST"),
			r.Method,
			r.RequestURI,
			r.Proto,
			r.Header["User-Agent"])

		next.ServeHTTP(w, r)
	})
}
