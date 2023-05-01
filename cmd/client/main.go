package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error to loading env: %s", err)
	}

	servAddr := os.Getenv("SERVER_ADDR")

	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		log.Fatalf("Connection error: %s", err)
	}
	log.Printf("Connection is starting to %s", servAddr)

	ch := make(chan string)
	defer close(ch)

	go readConsole(ch)
	go readFromServer(conn)

	for {
		val, ok := <-ch

		if ok {
			_, err := conn.Write([]byte(val))
			if err != nil {
				log.Fatalf("Error sending message: %s", err)
				break
			}
		} else {
			fmt.Println("Channel have not data")
		}
	}

	log.Println("Connection stopped")
	conn.Close()
}

func readConsole(ch chan string) {
	for {
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		message := line[:len(line)-1]
		ch <- message
	}
}

func readFromServer(conn net.Conn) {
	if conn == nil {
		panic("Connection is nil")
	}
	buf := make([]byte, 256)
	for {
		for i := 0; i < 256; i++ {
			buf[i] = 0
		}

		readed_len, _ := conn.Read(buf)
		if readed_len > 0 {
			fmt.Println(string(buf))
		}

	}
}
