package logger

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s(method)\t%s(uri)\t%s(name)\t%ss",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
