package server

import (
	nats "github.com/nats-io/go-nats"
	"log"

)

type ServerNATS struct {
	Conn nats.Conn
	close chan bool
}

func (s *ServerNATS) Stop() error {
	s.close <- true
	s.Conn.Close()
	return nil
}

func (s *ServerNATS) Start(handlers *Handlers) {
	s.close = make(chan bool)
	handleRequest := func(m *nats.Msg) {
		data := handleRequest(m.Data, handlers)
		log.Println("request: ", m.Data, "\n response: ", data)
		err := s.Conn.Publish(m.Reply, data)
		if err != nil {
			log.Println("Error accepting: ", err.Error())
		}
	}
	_, err := s.Conn.Subscribe("books", handleRequest)
	if err != nil {
		log.Fatal(err)
	}
	<- s.close
}

