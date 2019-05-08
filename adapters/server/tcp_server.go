package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type ServerTCP struct {
	Listener net.Listener
}


func (s *ServerTCP) Stop() error {
	return s.Listener.Close()
}

func (s *ServerTCP) Start(handlers *Handlers) {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			return
		}
		go newHandlerRequests(conn, handlers)
	}
}

func newHandlerRequests(conn net.Conn, handlers *Handlers) {
	defer conn.Close()
	for {
		buf, err := bufio.NewReader(conn).ReadSlice('\n')
		if err == io.EOF {
			fmt.Println("Client close connection.")
			return
		}
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			fmt.Println("Close connect.")
			return
		}
		_, _ = conn.Write(append(handleRequest(buf, handlers), byte('\n')))
	}
}
