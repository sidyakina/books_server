package infrastructure

import (
	"net"

	"github.com/sidyakina/books_server/adapters/server"
)

func InitServer(host, port string) (*server.Server, error) {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	return &server.Server{listener}, nil

}
