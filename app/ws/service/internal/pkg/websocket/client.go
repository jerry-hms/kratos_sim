package websocket

import (
	"sync"
)

type ClientMap map[int64]SessionID

func (c ClientMap) GetSessionId(id int64) SessionID {
	session_id, ets := c[id]
	if !ets {
		return ""
	}
	return session_id
}

type Client struct {
	*SessionManager
	clients ClientMap
	mtx     sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		clients: make(ClientMap),
	}
}

func (c *Client) Add(id int64, client_id SessionID) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.clients[id] = client_id
}

func (c *Client) Get(id int64) SessionID {
	return c.clients.GetSessionId(id)
}

func (c *Client) Remove(id int64) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	for curId := range c.clients {
		if curId == id {
			delete(c.clients, id)
		}
	}
}
