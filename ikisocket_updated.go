package ikisocket

import (
	"github.com/gofiber/websocket/v2"
)

func NewDefaultWebsocket() (kws *Websocket) {
	c := new(websocket.Conn)

	kws = &Websocket{
		ws: c,
		Locals: func(key string) interface{} {
			return c.Locals(key)
		},
		Params: func(key string, defaultValue ...string) string {
			return c.Params(key, defaultValue...)
		},
		Query: func(key string, defaultValue ...string) string {
			return c.Query(key, defaultValue...)
		},
		Cookies: func(key string, defaultValue ...string) string {
			return c.Cookies(key, defaultValue...)
		},
		queue:      make(chan message, 100),
		done:       make(chan struct{}, 1),
		attributes: make(map[string]interface{}),
		isAlive:    true,
	}

	// Generate uuid
	kws.UUID = kws.createUUID()

	// register the connection into the pool
	pool.set(kws)

	kws.fireEvent(EventConnect, nil, nil)

	// Run the loop for the given connection
	kws.run()
	return kws
}
