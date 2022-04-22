// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const sessionResourceURI = "/api/v1/session/"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	fromMqtt chan uint8

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// session resource url
	sessionResourceHost string
}

func newHub(apiHost string, apiPort int) *Hub {

	return &Hub{
		fromMqtt:   make(chan uint8),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),

		// For now, only http is supported
		sessionResourceHost: fmt.Sprintf("%s:%d", apiHost, apiPort),
	}
}

func (h *Hub) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			for client := range h.clients {
				close(client.send)
				delete(h.clients, client)
			}
			return
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.fromMqtt:
			wsMessage := []byte(fmt.Sprintf("%v", message))
			for client := range h.clients {
				select {
				case client.send <- wsMessage:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !h.authenticateClient(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: h, req: r, conn: conn, send: make(chan []byte, 2)}
	h.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func (h *Hub) authenticateClient(r *http.Request) bool {
	// Only http is supported for now
	newReq, err := http.NewRequestWithContext(r.Context(), "GET", fmt.Sprintf("http://%s%s/1/", h.sessionResourceHost, sessionResourceURI), nil)
	if err != nil {
		return false
	}

	// Copy original Host
	newReq.Host = r.Host

	// Copy cookies
	for _, cookie := range r.Cookies() {
		newReq.AddCookie(cookie)
	}

	client := &http.Client{}

	resp, err := client.Do(newReq)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return resp.StatusCode == 200
}
