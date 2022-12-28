package socketIO

import (
	socketio "github.com/googollee/go-socket.io"
	"sync"
)

type RealtimeEngine interface {
}

type rtEngine struct {
	server  *socketio.Server
	storage map[string][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{

		storage: make(map[string][]AppSocket),
		locker:  &sync.RWMutex{},
	}
}

func (engine *rtEngine) saveAppSocket(userName string, appSck AppSocket) {
	engine.locker.Lock()
	if v, ok := engine.storage[userName]; ok {
		engine.storage[userName] = append(v, appSck)
	} else {
		engine.storage[userName] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userName string) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()
	return engine.storage[userName]
}

func (engine *rtEngine) removeAppSocket(userName string, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()
	if v, ok := engine.storage[userName]; ok {
		for i, sck := range v {
			if sck == appSck {
				engine.storage[userName] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}
