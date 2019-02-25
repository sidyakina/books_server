package infrastructure

import (
	"github.com/sidyakina/books_server/adaptors/server"
	"net"
)

func InitServer(host, port string) (*server.Server, error){
	listener, err := net.Listen("tcp", host + ":" + port)
	if err != nil {
		return nil, err
	}
	return &server.Server{listener}, nil

}