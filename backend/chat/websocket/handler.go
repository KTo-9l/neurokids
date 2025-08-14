package main

import (
	"net"
	"strconv"

	"github.com/gobwas/ws"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) handleHttpToSocket(l logger.Logger, conn net.Conn) (err error) {
	var uid int
	u := ws.Upgrader{
		OnHeader: func(key, value []byte) error {
			l.Info(string(key), string(value))
			if string(key) == "X-User-Id" {
				uid, err = strconv.Atoi(string(value))
				if err != nil {
					l.Error("Parsing x-user-id", err)
					return err
				}
			}
			return nil
		},
	}

	_, err = u.Upgrade(conn)
	if err != nil {
		l.Error("Upgrade To Websocket", err)
		return
	}

	l.Info("Upgraded for client", strconv.Itoa(uid))

	go addClient(l, conn, uid, s.pubs_getter)
	return nil
}

func (s *service) Close(l logger.Logger) error {
	return nil
}
