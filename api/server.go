package api

import (
	"log"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/api/key", s.HandleGetKey)
	log.Println("Server listening on address ", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) HandleGetKey(w http.ResponseWriter, r *http.Request) {

}
