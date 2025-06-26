# 💬 PP - WebSocket Chat App

A simple real-time chat application built in Go using **WebSockets**, designed as a learning project to explore how persistent two-way communication works between a frontend and backend server.

---

## 🚀 Overview

This project is a **personal experiment** and learning exercise to deeply understand and implement:

- 🔌 WebSocket connections in Go using `github.com/gorilla/websocket`
- 💡 Real-time client/server communication
- 🧠 Managing connected clients and chatrooms with concurrency
- 🧱 Clean architectural separation (Manager, Client, Shared Interfaces)
- 🧪 HTML + Vanilla JS frontend for testing and interacting with the server

---

## 📁 Project Structure

```text
web-socket-chatApp/
├── cmd/
│   └── main.go                # Entry point for the server
├── client/
│   └── client.go              # WebSocket client abstraction
├── manager/
│   └── manager.go             # Manages active clients and broadcasts
├── shared/
│   └── shared.go              # Shared interfaces to avoid import cycles
├── frontend/
│   └── index.html             # Simple HTML + JS UI to test the chat
└── README.md
```
