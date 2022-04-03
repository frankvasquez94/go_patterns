package main

//

import (
	"fmt"
	"time"
)

const PatternName = "Option pattern design"

// Main structure for a server
// port is the unique required value
// timeOut and maxConnection are optional
type Server struct {
	port           string
	timeOut        time.Duration
	maxConnections int
}

// Option Define a type Option that is
// a function wich takes a Server pointer
type Option func(*Server)

func WithTimeOut(timeOut time.Duration) Option {
	return func(s *Server) {
		s.timeOut = timeOut
	}
}

func WithMaxConnections(maxConnection int) Option {
	return func(s *Server) {
		s.maxConnections = maxConnection
	}
}

func (s *Server) ToString() string {
	resp := "Port: " + s.port

	if s.timeOut != 0 {
		resp += ", time out: " + fmt.Sprintf("%v", s.timeOut)
	}

	if s.maxConnections != 0 {
		resp += ", max connections: " + fmt.Sprintf("%v", s.maxConnections)
	}

	return resp
}

func NewServer(port string, options ...Option) *Server {
	server := &Server{
		port: port,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func main() {
	p := fmt.Println
	p(PatternName)

	p("=========== Basic server =============")

	basicServer := NewServer("90")
	p(basicServer.ToString())

	p("======== Server with optional timeout ===========")

	withTimeOutServer := NewServer("90", WithTimeOut(3600))
	p(withTimeOutServer.ToString())

	p("============== Server with timeout and max connections")

	withTimeOutAndMaxConnServer := NewServer("90", WithTimeOut(3600), WithMaxConnections(50))
	p(withTimeOutAndMaxConnServer.ToString())

}
