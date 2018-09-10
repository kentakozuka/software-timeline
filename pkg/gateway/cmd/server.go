package cmd

import (
	"log"
	"net/http"
)

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
		server: &http.Server{
			Addr: ":8000",
		},
	}
}

type Server struct {
	server *http.Server
	mux    *http.ServeMux
}

func (s *Server) ListenAndServe() {
	s.routes()
	s.server.Handler = s.mux
	log.Fatal(s.server.ListenAndServeTLS("cert/server.crt", "cert/server.key"))
}

func (s *Server) routes() {
	s.mux.HandleFunc("/hello", s.handleHello())
	s.mux.HandleFunc("/timeline", s.handleTimeline())
}

func (s *Server) handleHello() http.HandlerFunc {
	// you can prepare something.
	log.Printf("hello handler is registered.")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got connection: %s", r.Proto)
		w.Write([]byte("Hello"))
	}
}

func (s *Server) handleTimeline() http.HandlerFunc {
	// you can prepare something.
	log.Printf("timeline handler is registered.")
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			w.Write([]byte("timeline get"))

		case http.MethodPost:
			w.Write([]byte("timeline post"))

		case http.MethodPut:
			w.Write([]byte("timeline put"))

		case http.MethodDelete:
			w.Write([]byte("timeline delete"))

		default:
			w.Write([]byte("timeline error"))
		}
	}
}
