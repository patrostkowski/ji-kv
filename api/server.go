package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
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
	log.Println("API server listening on address ", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) HandleKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//TODO find better way than getting size 1 lists
	queryKey := r.URL.Query()["key"]
	if len(queryKey) == 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	switch r.Method {
	case http.MethodGet:

		time.Sleep(800 * time.Millisecond)

		//TODO Read from blob
		value := map[string]string{
			"key":   queryKey[0],
			"value": strconv.Itoa(rand.Int())}

		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(value)
		if err != nil {
			log.Fatalln("There was an error decoding the request body into the struct")
		}

		return
	case http.MethodPost:

		time.Sleep(1422 * time.Millisecond)

		queryValue := r.URL.Query()["value"]
		if len(queryKey) == 0 {
			err := json.NewEncoder(w).Encode(http.StatusBadRequest)
			if err != nil {
				log.Fatalln("There was an error decoding the request body into the struct")
			}
			return
		}

		//TODO Write to blob
		value := map[string]string{
			"key":   queryKey[0],
			"value": queryValue[0]}

		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(value)
		if err != nil {
			log.Fatalln("There was an error decoding the request body into the struct")
		}
		return

	case http.MethodDelete:

		time.Sleep(800 * time.Millisecond)

		//TODO delete from blob
		value := map[string]string{
			"key":   queryKey[0],
			"value": strconv.Itoa(rand.Int())}

		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(value)
		if err != nil {
			log.Fatalln("There was an error decoding the request body into the struct")
		}
		return

	default:
		time.Sleep(2468 * time.Millisecond)
		json.NewEncoder(w).Encode(http.StatusForbidden)
		return
	}
}
