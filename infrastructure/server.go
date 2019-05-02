package infrastructure

import (
	"net"

	"github.com/sidyakina/books_server/adapters/server"
)

func InitServer(port string) (*server.Server, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	return &server.Server{listener}, nil

}
