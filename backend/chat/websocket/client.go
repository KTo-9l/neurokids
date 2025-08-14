package main

import (
	"context"
	"fmt"
	"net"
	"slices"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/okonma-violet/services/logs/logger"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type Hub struct {
	mu      sync.Mutex
	clients map[int][]*Client
}

type Client struct {
	conn   net.Conn
	ctx    context.Context
	cancel context.CancelFunc
	id     int
}

var hub *Hub

func addClient(l logger.Logger, conn net.Conn, id int, pubs universalservice_nonepoll.Publishers_getter) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:   conn,
		ctx:    ctx,
		cancel: cancel,
		id:     id,
	}

	if hub.clients[id] == nil {
		hub.clients[id] = make([]*Client, 0, 10)
	}
	hub.clients[id] = append(hub.clients[id], client)

	go client.listen(l, pubs)
	go client.ping(l)
}

func (client *Client) listen(l logger.Logger, pubs universalservice_nonepoll.Publishers_getter) {
	for {
		frame, err := ws.ReadFrame(client.conn)
		if err != nil {
			l.Error("Error reading frame:", err)
			return
		}

		if frame.Header.Masked {
			ws.Cipher(frame.Payload, frame.Header.Mask, 0)
			frame.Header.Masked = false
		}

		switch frame.Header.OpCode {
		case ws.OpClose:
			l.Info(fmt.Sprint(client.id), "Client closed connection")
			client.kill(l)
			return
		case ws.OpPing:
			l.Info(fmt.Sprint(client.id), "Received Ping")
			err = ws.WriteFrame(client.conn, ws.NewPongFrame(frame.Payload))
		case ws.OpPong:
			l.Info(fmt.Sprint(client.id), "Received Pong")
		case ws.OpText:
			l.Info(fmt.Sprint(client.id), "Received message")
			err = unmarshalMessage(l, frame.Payload, pubs)
		case ws.OpBinary:
			l.Info(fmt.Sprint(client.id), "Received binary message")
			err = ws.WriteFrame(client.conn, ws.NewBinaryFrame([]byte("Echo: "+string(frame.Payload))))
		default:
			l.Warning(fmt.Sprint(client.id), string(frame.Header.OpCode))
		}

		if err != nil {
			l.Error("Answer", err)
		}
	}
}

func (client *Client) ping(l logger.Logger) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-client.ctx.Done():
			l.Debug(fmt.Sprint(client.id), "Context done")
			return
		case <-ticker.C:
			err := ws.WriteFrame(client.conn, ws.NewPingFrame([]byte("Hey?")))
			if err != nil {
				l.Error(fmt.Sprintf("%s] [%s", client.id, "WriteFrame"), err)
				client.kill(l)
				return
			}
		}
	}
}

func (client *Client) kill(l logger.Logger) {
	client.cancel()

	err := client.conn.Close()
	if err != nil {
		l.Error("Closing connection", err)
	}

	hub.mu.Lock()
	defer hub.mu.Unlock()

	clientIndex := slices.Index(hub.clients[client.id], client)
	hub.clients[client.id] = slices.Delete(hub.clients[client.id], clientIndex, clientIndex+1)

	if len(hub.clients[client.id]) == 0 {
		delete(hub.clients, client.id)
	}
}
