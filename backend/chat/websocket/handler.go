package main

import (
	"net"

	"dikobra3/mongoApi"

	"github.com/big-larry/mgo/bson"
	"github.com/gobwas/ws"

	"github.com/okonma-violet/services/logs/logger"
)

func (s *service) handleHttpToSocket(l logger.Logger, conn net.Conn) (err error) {
	var uid bson.ObjectId
	u := ws.Upgrader{
		OnHeader: func(key, value []byte) error {
			l.Info(string(key), string(value))
			if string(key) == "X-User-Id" {
				uid = mongoApi.StringToObjectId(string(value))
			}
			return nil
		},
	}

	_, err = u.Upgrade(conn)
	if err != nil {
		l.Error("Upgrade To Websocket", err)
		return
	}

	l.Info("Upgraded for client", string(uid))

	go addClient(l, conn, uid, s.pubs_getter)
	return nil
}

func (s *service) Close(l logger.Logger) error {
	return nil
}
