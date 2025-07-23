package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/big-larry/suckhttp"
	"github.com/gobwas/ws"
	"github.com/okonma-violet/services/logs/logger"
	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type WsMessage struct {
	Event string          `json:"event"`
	From  string          `json:"from"`
	To    []string        `json:"to"`
	Data  json.RawMessage `json:"data"`
}

func unmarshalMessage(l logger.Logger, payload []byte, pubs universalservice_nonepoll.Publishers_getter) (err error) {
	var message *WsMessage
	err = json.Unmarshal(payload, &message)
	if err != nil {
		l.Error("Unmarshal message", err)
		return
	}

	l.Debug("Unmarshalled data", fmt.Sprint(string(message.Data)))

	err = message.handle(l, pubs)
	if err != nil {
		l.Error("Handle message", err)
		return
	}

	return nil
}

func (message *WsMessage) handle(l logger.Logger, pubs universalservice_nonepoll.Publishers_getter) (err error) {
	var publisher *universalservice_nonepoll.Publisher

	switch message.Event {
	case "createChat":
		publisher = pubs.Get(chat_create_chat)
	case "updateChat":
		publisher = pubs.Get(chat_update_chat)
	case "addMessage":
		publisher = pubs.Get(chat_add_message)
	default:
		return errors.New("Unknown event")
	}

	resp, err := message.send(l, publisher)
	if err != nil {
		l.Error("HandleMessage", err)
		return err
	}

	newMessage := &WsMessage{
		Event: message.Event,
		From:  message.From,
		To:    message.To,
		Data:  resp.GetBody(),
	}

	err = newMessage.sendToClients(l)
	if err != nil {
		l.Error("Send Message To Clients", err)
		return
	}
	return nil
}

func (message *WsMessage) send(l logger.Logger, pub_getter *universalservice_nonepoll.Publisher) (response *suckhttp.Response, err error) {
	request, err := universalservice_nonepoll.CreateHTTPRequest("POST")
	if err != nil {
		l.Error("WsMessage.send", err)
		return
	}
	request.Body = message.Data
	request.AddHeader("Content-Type", "application/json")

	resp, err := pub_getter.SendHTTP(request)
	return resp, err
}

func (message *WsMessage) sendToClients(l logger.Logger) (err error) {
	bts, err := json.Marshal(message)
	if err != nil {
		l.Error("Marshal message", err)
		return
	}

	for _, uid := range message.To {
		if hub.clients[uid] != nil && len(hub.clients[uid]) != 0 {
			for _, client := range hub.clients[uid] {
				err = ws.WriteFrame(client.conn, ws.NewTextFrame(bts))
				if err != nil {
					l.Error("Sending message to client", err)
					return
				}
			}
		}
	}
	return nil
}
