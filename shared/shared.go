package shared

import "github.com/gorilla/websocket"

// ManagerInterface defines what a manager must do (for now or future expansion)
type ManagerInterface interface {
	AddClient(client ClientInterface)
	RemoveClient(client ClientInterface)
}

// ClientInterface defines minimal fields needed for cleanup
type ClientInterface interface {
	Connection() *websocket.Conn
}
