package proto

import (
	"gitlab.com/hasura/hasuractl-go/pkg/ngrok/src/ngrok/conn"
)

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
