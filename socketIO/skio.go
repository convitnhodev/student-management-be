package socketIO

import (
	_const "managerstudent/common/const"
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	// ID returns the connection id.
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	Context() interface{}
	SetContext(context interface{})
	Namespace() string
	Emit(msg string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
	_const.Requester
}

type appSocket struct {
	Conn
	_const.Requester
}

func NewAppSocket(conn Conn, requester _const.Requester) *appSocket {
	return &appSocket{
		Conn:      conn,
		Requester: requester,
	}
}
