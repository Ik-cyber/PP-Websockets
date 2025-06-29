package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ik-cyber/web-socket-chatApp/manager"
)

func main() {
	setupApi()

	fmt.Println("✅ Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupApi() {
	manager := manager.NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.ServeWS)
}
