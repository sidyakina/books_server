package infrastructure

import (
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
