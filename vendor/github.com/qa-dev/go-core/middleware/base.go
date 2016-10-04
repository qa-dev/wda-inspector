package middleware

import (
	"github.com/qa-dev/go-core/log"
	"net/http"
	"runtime/debug"
)

func New(h http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal("Panic: %+v\n%s", err, debug.Stack())
			}
		}()

		lrw := &LoggedResponseWriter{responseWriter: resp}

		h.ServeHTTP(lrw, req)
		log.Info("%v %v %v (%v)",
			lrw.Status(), req.Method, req.URL.Path, req.RemoteAddr)
	})
}
