package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sidyakina/books_server/connections/postgres"
	"github.com/sidyakina/books_server/use_case"
	"io"
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

		request := make(map[string]interface{})
		err = json.Unmarshal(buf, &request)
		if err != nil {
			conn.Write(use_case.ErrorResult("can't unmarshal request"))
			continue
		}
		fmt.Printf("get request %v\n", request)
		cmd, _ := request["cmd"]
		var response []byte
		switch cmd {
		case "getAllBooks":
			response = use_case.GetAllBooks(pg)
		case "addBook":
			response = use_case.AddBook(pg, buf)
		case "deleteBook":
			response = use_case.DeleteBook(pg, buf)
		default:
			response = use_case.ErrorResult(fmt.Sprintf("cmd: %v is not defined", cmd))
		}
		//fmt.Printf("send response %v\n", response)
		_, _  = conn.Write(response)
	}


}