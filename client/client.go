package client

import (
	"log"

	"github.com/Ik-cyber/web-socket-chatApp/shared"
	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	// the websocket connection
	connection *websocket.Conn

	// manager is the manager used to manage the client but implements the ManagerInterface
	manager shared.ManagerInterface
}

// NewClient is used to initialize a new Client with all required values initialized
func NewClient(conn *websocket.Conn, manager shared.ManagerInterface) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

// This is suppose to be ran as a goroutine
func (c *Client) ReadMessages() {
	defer func() {
		// Graceful Close the Connection once this
		// function is done
		c.manager.RemoveClient(c)
	}()
	for {
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil {
			// If Connection is closed, we will Recieve an error here
			// We only want to log Strange errors, but not simple Disconnection
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}
		log.Println("MessageType: ", messageType)
		log.Println("Payload: ", string(payload))
	}
}

func (c *Client) Connection() *websocket.Conn {
	return c.connection
}
