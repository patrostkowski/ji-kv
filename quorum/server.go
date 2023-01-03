package quorum

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type Server struct {
	listenAddr     string
	ln             net.Listener
	quitch         chan struct{}
	quorumEntities []Entity
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	log.Println("Quorum server listening on address ", s.listenAddr)

	defer ln.Close()
	s.ln = ln

	go s.acceptConn()

	<-s.quitch
	return nil
}

func (s *Server) acceptConn() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			continue
		}

		log.Printf("Accepted connection from %v", conn.RemoteAddr())
		e := Entity{
			name: strconv.Itoa(rand.Int()),
			addr: conn.RemoteAddr().String(),
		}

		s.quorumEntities = append(s.quorumEntities, e)

		go s.readLoop(conn)
		go s.refreshEntities()
	}
}

func (*Server) closeLoop(conn net.Conn) {
	log.Println("Closing conn", conn.RemoteAddr())
	conn.Close()
}

func (*Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			continue
		}

		msg := buffer[:n]
		log.Printf("Recived message from (%v): %v", conn.RemoteAddr().String(), string(msg))
	}
}

func (s *Server) refreshEntities() {
	for {
		log.Println("Quorum entities", s.quorumEntities)
		time.Sleep(1 * time.Second)
	}
}
