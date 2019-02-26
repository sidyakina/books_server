package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/sidyakina/books_server/domain"
)

type Server struct {
	Listener net.Listener
}

type handlerGet interface {
	GetAllBooks() ([]domain.Book, string)
}

type handlerAdd interface {
	AddBook(request domain.RequestAdd) (int32, string)
}

type handlerRemove interface {
	RemoveBook(request domain.RequestRemove) (int32, string)
}

type Handlers struct {
	get    handlerGet
	add    handlerAdd
	remove handlerRemove
}

func InitHandlers(hget handlerGet, hadd handlerAdd, hremove handlerRemove) *Handlers {
	return &Handlers{get: hget, add: hadd, remove: hremove}
}

func (s *Server) Stop() error {
	return s.Listener.Close()
}

func (s *Server) Start(handlers *Handlers) {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
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
		_, _ = conn.Write(handleRequest(buf, handlers))
	}
}

func handleRequest(buf []byte, handlers *Handlers) []byte {
	request := make(map[string]interface{})
	result := make(map[string]interface{})
	err := json.Unmarshal(buf, &request)
	if err != nil {
		fmt.Println(err)
		result["error"] = "error: can't unmarshal request."
		response, _ := json.Marshal(result)
		return append(response, byte('\n'))
	}
	fmt.Printf("get request %v\n", request)
	cmd, _ := request["cmd"]
	result["cmd"] = cmd
	switch cmd {
	case "getAllBooks":
		books, err := handlers.get.GetAllBooks()
		result["books"] = books
		result["error"] = err
	case "addBook":
		request := domain.RequestAdd{}
		err := json.Unmarshal(buf, &request)
		if err != nil {
			fmt.Print(err)
			result["id"] = 0
			result["error"] = "error: can't unmarshal request."
		} else {
			id, err := handlers.add.AddBook(request)
			result["id"] = id
			result["error"] = err
		}
	case "deleteBook":
		request := domain.RequestRemove{}
		err := json.Unmarshal(buf, &request)
		if err != nil {
			fmt.Println(err)
			result["error"] = "error while deleteBook: cant't unmarshal Params"
		} else {
			id, err := handlers.remove.RemoveBook(request)
			result["id"] = id
			result["error"] = err
		}
	default:
		result["error"] = fmt.Sprintf("cmd: %v is not defined", cmd)
	}
	//fmt.Printf("send response %v\n", response)
	response, _ := json.Marshal(result)
	return append(response, byte('\n'))
}
