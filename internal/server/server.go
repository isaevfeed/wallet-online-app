package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"wallet/internal/client"
)

var response *client.ClientResponse

type Server struct {
	addr        string
	connections map[int]*client.Connection
}

func NewServer(addr string) *Server {
	return &Server{
		addr:        addr,
		connections: make(map[int]*client.Connection),
	}
}

func (s *Server) Start() {
	serv, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Server starting error: %s", err)
	}
	log.Printf("Server is starting on %s", s.addr)

	i := 0

	response = &client.ClientResponse{}

	for {
		conn, err := serv.Accept()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalf("Connection refused error: %s", err)
				break
			}
		}

		client := client.NewConnection(conn, i)
		s.connections[i] = client

		log.Println("New connection: ", client.Name)

		for key := range s.connections {
			sibConn := s.connections[key].Conn
			sibConn.Write([]byte("New client: " + client.Name))
		}

		go s.connectionProcess(i)
		i++
	}

	log.Println("Server stopped")
}

func (s *Server) connectionProcess(i int) {
	client := s.connections[i]
	conn := client.Conn
	buf := make([]byte, 256)
	defer conn.Close()

	response.Data = s.connections

	for {
		read_len, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Connection process error: %s", err)
				delete(s.connections, i)
				break
			}
		}

		message := string(buf[:read_len])
		fmt.Println(message)

		clients, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error marshal object %s", err)
		}

		//take>1

		var action string
		var idx int
		items := strings.Split(message, ">")
		if len(items) > 1 {
			action = items[0]
			items[1] = items[1][:len(items[1])-1]
			idx, err = strconv.Atoi(items[1])
			if err != nil {
				log.Fatal("Index is not a number %s", err)
			}
		}

		guestClient := s.connections[idx]

		if client != nil && idx != i {
			if action == "take" {
				guestClient.Wall.Take(client.Wall, 100)

				for key := range s.connections {
					sibConn := s.connections[key].Conn
					sibConn.Write([]byte("User #" + strconv.Itoa(i) + " take 100 from #" + strconv.Itoa(idx)))
				}

				continue
			}
		} else {
			conn.Write(clients)

			for key := range s.connections {
				sibConn := s.connections[key].Conn
				sibConn.Write([]byte(message))
			}
		}
	}
}
