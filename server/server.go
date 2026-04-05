package server

/*
Create your own text based protocol over TCP/IP to connect with NATS server for clients.
We can use websocket over TCP/IP to connect server to client and vice versa
*/

// Config stores configuration of the server
type Config struct {
}

type Server struct {
	cfg *Config
}

// Start starts the server
func (s *Server) Start() {

}
