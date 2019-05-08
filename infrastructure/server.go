package infrastructure

import (
	nats "github.com/nats-io/go-nats"
	"github.com/sidyakina/books_server/adapters/server"
	"net"
)

func InitServerTCP(port string) (*server.ServerTCP, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	return &server.ServerTCP{listener}, nil

}

func InitServerNATS(host, port string) (*server.ServerNATS, error) {
	nc, err := nats.Connect(host + ":" + port)
	if err != nil {
		return nil, err
	}
	return &server.ServerNATS{Conn: *nc}, nil

}