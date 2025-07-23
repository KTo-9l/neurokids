package main

import (
	"context"

	"github.com/okonma-violet/services/universalservice_nonepoll"
)

type config struct {
}

type service struct {
	pubs_getter universalservice_nonepoll.Publishers_getter
}

const (
	chat_get_chats    = "chat_get_chats"
	chat_create_chat  = "chat_create_chat"
	chat_update_chat  = "chat_update_chat"
	chat_get_messages = "chat_all_messages"
	chat_add_message  = "chat_add_message"
)

func (c *config) InitFlags() {}

func (c *config) PrepareHandling(ctx context.Context, pubs_getter universalservice_nonepoll.Publishers_getter) (universalservice_nonepoll.BaseHandleFunc, universalservice_nonepoll.Closer, error) {
	s := &service{
		pubs_getter: pubs_getter,
	}

	initHub()

	return s.handleHttpToSocket, s, nil
}

func main() {
	services := []universalservice_nonepoll.ServiceName{
		chat_create_chat,
		// chat_get_all_chats,
		chat_update_chat,
		// chat_get_all_messages,
		chat_add_message,
	}
	universalservice_nonepoll.InitNewServiceWithoutName(&config{}, 1, services...)
}
