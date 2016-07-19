package rtm

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	StateClosed = "closed"
	StateOpen   = "open"
)

var (
	ErrLoopClosed = errors.New("loop closed")
)

type Loop struct {
	lock   sync.RWMutex
	state  string
	wsHost string
	conn   *websocket.Conn
	callId uint64

	C chan Message
}

func NewLoop(wsHost string, backlog int) *Loop {
	return &Loop{
		lock: sync.RWMutex{},

		state:  StateClosed,
		wsHost: wsHost,
		C:      make(chan Message, backlog),
	}
}

func (l *Loop) Start() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	conn, _, err := websocket.DefaultDialer.Dial(l.wsHost, nil)
	if err != nil {
		return err
	}

	l.conn = conn
	l.state = StateOpen

	go l.readMessage()

	return nil
}

func (l *Loop) Close() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.conn.Close()
	l.state = StateClosed
}

func (l *Loop) State() string {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.state
}

func (l *Loop) Ping() error {
	return l.writeMessage(Message{"type": MessageTypePing})
}

func (l *Loop) Keepalive(interval *time.Ticker) error {
	for {
		select {
		case <-interval.C:
			if err := l.Ping(); err != nil {
				return err
			}
		}
	}
}

func (l *Loop) readMessage() {
	for {
		if l.State() == StateClosed {
			return
		}

		_, raw, err := l.conn.ReadMessage()
		if err != nil {
			log.Printf("rtm loop read: %s", err)
			return
		}

		message := Message{}
		if err = json.Unmarshal(raw, &message); err != nil {
			log.Printf("rtm loop parse: %s %s", err, raw)
			continue
		}

		l.C <- message
	}
}

func (l *Loop) writeMessage(m Message) error {
	if l.State() == StateClosed {
		return ErrLoopClosed
	}

	if _, present := m["call_id"]; !present {
		m["call_id"] = l.nextCallId()
	}

	raw, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return l.conn.WriteMessage(websocket.TextMessage, raw)
}

// TODO use sync.atomic
func (l *Loop) nextCallId() uint64 {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.callId = l.callId + 1
	return l.callId
}
