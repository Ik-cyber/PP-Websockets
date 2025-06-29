package manager

import (
	"log"
	"net/http"
	"sync"

	"github.com/Ik-cyber/web-socket-chatApp/client"
	"github.com/Ik-cyber/web-socket-chatApp/shared"
	"github.com/gorilla/websocket"
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Manager struct {
	clients map[shared.ClientInterface]bool

	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[shared.ClientInterface]bool),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New connection")
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := client.NewClient(conn, m) // `m` is *Manager which implements ManagerInterface

	m.AddClient(c) // works because c implements ClientInterface

	// Start the client's message reading loop in a go routine
	go c.ReadMessages()
}

func (m *Manager) AddClient(client shared.ClientInterface) {
	// Lock to aviod data race
	m.Lock()
	defer m.Unlock()

	m.clients[client] = true
}

// removeClient will remove the client and clean up
func (m *Manager) RemoveClient(client shared.ClientInterface) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.clients[client]; ok {
		client.Connection().Close()
		delete(m.clients, client)
	}
}
