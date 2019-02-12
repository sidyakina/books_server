package server

import (
	"encoding/json"
	"fmt"
	"github.com/sidyakina/books_server/connections/postgres"
	"github.com/sidyakina/books_server/use_case"
	"net"
)

type Server struct {
	pg *postgres.ConnectPG
	listener net.Listener
}

func Init(pg *postgres.ConnectPG, host, port string) (*Server, error){
	listener, err := net.Listen("tcp", host + ":" + port)
	if err != nil {
		return nil, err
	}
	return &Server{pg, listener}, nil

}

func (s *Server) Stop () error{
	return s.listener.Close()
}

func (s *Server)Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			//os.Exit(1)
		}
		go handleRequest(conn, s.pg)
	}
}


func handleRequest(conn net.Conn, pg *postgres.ConnectPG) {
	defer conn.Close()
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	request := make(map[string]interface{})
	err = json.Unmarshal(buf[:reqLen], &request)
	if err != nil {
		conn.Write(use_case.ErrorResult("can't unmarshal request"))
		return
	}
	fmt.Printf("get request %v\n", request)
	cmd, _ := request["cmd"]
	var response []byte
	switch cmd {
	case "getAllBooks":
		response = use_case.GetAllBooks(pg)
	case "addBook":
		response = use_case.AddBook(pg, request)
	case "deleteBook":
		response = use_case.DeleteBook(pg, request)
	default:
		response = use_case.ErrorResult(fmt.Sprintf("cmd: %v is not defined", cmd))
	}
	_, _  = conn.Write(response)
}