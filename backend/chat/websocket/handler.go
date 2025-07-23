package main

import (
	"net"

	"github.com/gobwas/ws"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) handleHttpToSocket(l logger.Logger, conn net.Conn) (err error) {
	var uid string
	u := ws.Upgrader{
		OnHeader: func(key, value []byte) error {
			if string(key) == "X-User-Id" {
				uid = string(value)
			}
			return nil
		},
	}

	_, err = u.Upgrade(conn)
	if err != nil {
		l.Error("Upgrade To Websocket", err)
		return
	}

	l.Info("Upgraded for client", uid)

	go addClient(l, conn, uid, s.pubs_getter)
	return nil
}

func (s *service) Close(l logger.Logger) error {
	return nil
}
