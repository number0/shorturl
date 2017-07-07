package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/julienschmidt/httprouter"
	"go.iondynamics.net/templice"
)

// Server ...
type Server struct {
	bind      string
	templates *templice.Template
	router    *httprouter.Router
}

func (s *Server) render(w http.ResponseWriter, tmpl string, data interface{}) {
	err := s.templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// IndexHandler ...
func (s *Server) IndexHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s.render(w, "index", nil)
	}
}

// ShortenHandler ...
func (s *Server) ShortenHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		url := r.Form.Get("url")
		u, err := NewURL(url)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}

		s.render(w, "url", struct{ ID string }{ID: u.ID()})
	}
}

// RedirectHandler ...
func (s *Server) RedirectHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		if id == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		u, ok := LookupURL(id)
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, u.url, http.StatusFound)
	}
}

// ListenAndServe ...
func (s *Server) ListenAndServe() {
	log.Fatal(http.ListenAndServe(s.bind, s.router))
}

func (s *Server) initRoutes() {
	s.router.GET("/", s.IndexHandler())
	s.router.POST("/", s.ShortenHandler())
	s.router.GET("/:id", s.RedirectHandler())
}

// NewServer ...
func NewServer(bind string) *Server {
	server := &Server{
		bind:      bind,
		router:    httprouter.New(),
		templates: templice.New(rice.MustFindBox("templates")),
	}

	err := server.templates.Load()
	if err != nil {
		log.Panicf("error loading templates: %s", err)
	}

	server.initRoutes()

	return server
}
