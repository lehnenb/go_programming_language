package client

import (
	"fmt"
	"log"
	"net/http"
)

// ServeClient handles websocket requests from the peer.
func ServeClient(broadcast genericBroadcast, w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)

  if err != nil {
    log.Println(err)
    return
  }

  client := &Client{broadcast: broadcast, conn: conn}
  err = client.broadcast.RegisterClient(client)

  if err != nil {
    fmt.Println(err)
    client.broadcast.UnregisterClient(client)
    client.conn.Close()
    return
  }

  // Allow collection of memory referenced by the caller by doing all work in
  // new goroutines.
  go client.writePump()
  go client.readPump()
}
