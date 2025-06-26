package manager

import (
	"log"
	"net/http"
	"sync"

	"github.com/Ik-cyber/web-socket-chatApp/client"
	"github.com/Ik-cyber/web-socket-chatApp/shared"
	"github.com/gorilla/websocket"
)

/*
websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
*/
var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Manager is used to hold references to all Clients Registered, and Broadcasting etc
type Manager struct {
	clients map[shared.ClientInterface]bool

	sync.RWMutex
}

// NewManager is used to initalize all the values inside the manager
func NewManager() *Manager {
	return &Manager{
		clients: make(map[shared.ClientInterface]bool),
	}
}

// serveWS is a HTTP Handler that the has the Manager that allows connections
func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")
	// Begin by upgrading the HTTP request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := client.NewClient(conn, m) // `m` is *Manager which implements ManagerInterface

	// Add the client to the manager
	m.AddClient(c) // works because c implements ClientInterface

	// Start the client's message reading loop
	go c.ReadMessages()
}

func (m *Manager) AddClient(client shared.ClientInterface) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	m.clients[client] = true
}

// removeClient will remove the client and clean up
func (m *Manager) RemoveClient(client shared.ClientInterface) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[client]; ok {
		// close connection
		client.Connection().Close()
		// remove
		delete(m.clients, client)
	}
}
