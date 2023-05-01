package client

import (
	"net"
	"strconv"

	"wallet/internal/wallet"
)

type ClientResponse struct {
	Data map[int]*Connection `json:"data"`
}

type Connection struct {
	Conn   net.Conn
	Name   string         `json:"name"`
	Number int            `json:"number"`
	Wall   *wallet.Wallet `json:"wallet"`
}

func NewConnection(conn net.Conn, number int) *Connection {
	return &Connection{
		Conn:   conn,
		Name:   "Client #" + strconv.Itoa(number+1),
		Number: number,
		Wall:   wallet.NewWallet(),
	}
}
