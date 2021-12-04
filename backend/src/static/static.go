package static

import (
	"log"
	"net/http"
)

type Static struct {
	Handler http.HandlerFunc
}

func New(path string) Static {
	return Static{
		Handler: loggingHandler(http.FileServer(http.Dir(path))).ServeHTTP,
	}
}

func loggingHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("STATIC:%s:%s\n", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
