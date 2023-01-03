package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	http.HandleFunc("/v1/api", s.HandleKey)
	log.Println("Server listening on address ", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) HandleKey(w http.ResponseWriter, r *http.Request) {
	//TODO find better way than getting size 1 lists
	queryKey := r.URL.Query()["key"]
	if len(queryKey) == 0 {
		return
	}
	switch r.Method {
	case http.MethodGet:
		//TODO Read from blob
		value := map[string]string{
			"status": strconv.Itoa(http.StatusOK),
			"key":    queryKey[0],
			"value":  strconv.Itoa(rand.Int())}

		json.NewEncoder(w).Encode(value)
		return
	case http.MethodPost:

		queryValue := r.URL.Query()["value"]
		if len(queryKey) == 0 {
			json.NewEncoder(w).Encode(http.StatusBadRequest)
			return
		}

		//TODO Write to blob
		value := map[string]string{
			"status": strconv.Itoa(http.StatusOK),
			"key":    queryKey[0],
			"value":  queryValue[0]}

		json.NewEncoder(w).Encode(value)
		return

	case http.MethodDelete:

		//TODO delete from blob
		value := map[string]string{
			"status": strconv.Itoa(http.StatusOK),
			"key":    queryKey[0],
			"value":  strconv.Itoa(rand.Int())}

		json.NewEncoder(w).Encode(value)
		return

	default:
		json.NewEncoder(w).Encode(http.StatusForbidden)
		return
	}
}
