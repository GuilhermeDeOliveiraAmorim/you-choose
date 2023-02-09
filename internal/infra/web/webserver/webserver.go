package webserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
	Verb          string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, verb string) {
	s.Handlers[path] = handler
	s.Verb = verb
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		switch s.Verb {
		case "Get":
			s.Router.Get(path, handler)
		case "Delete":
			s.Router.Delete(path, handler)
		case "Patch":
			s.Router.Patch(path, handler)
		case "Put":
			s.Router.Put(path, handler)
		case "Post":
			s.Router.Post(path, handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
