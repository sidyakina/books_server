package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sidyakina/books_server/use_case"
	"io"
	"net"
)

type Server struct {
	Listener net.Listener
}

type handlerGet interface {
	GetAllBooks() []byte
}

type handlerAdd interface {
	AddBook (request []byte) []byte
}

type handlerRemove interface {
	RemoveBook (request [] byte) []byte
}

type Handlers struct {
	get handlerGet
	add handlerAdd
	remove handlerRemove
}

func InitHandlers(hget handlerGet, hadd handlerAdd, hremove handlerRemove) *Handlers {
	return &Handlers{get: hget, add: hadd, remove: hremove}
}

func (s *Server) Stop () error{
	return s.Listener.Close()
}

func (s *Server)Start(handlers *Handlers) {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn, handlers)
	}
}


func handleRequest(conn net.Conn, handlers *Handlers) {
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
			response = handlers.get.GetAllBooks()
		case "addBook":

			response = handlers.add.AddBook(buf)
		case "deleteBook":
			response = handlers.remove.RemoveBook(buf)
		default:
			response = use_case.ErrorResult(fmt.Sprintf("cmd: %v is not defined", cmd))
		}
		//fmt.Printf("send response %v\n", response)
		_, _  = conn.Write(response)
	}


}