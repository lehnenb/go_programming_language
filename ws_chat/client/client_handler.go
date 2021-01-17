package client 

import (
  "net/http"
  "log"
  "github.com/lehnenb/go_programming_language/ws_chat/broadcast"
)

// serveWs handles websocket requests from the peer.
func serveWs(broadcast *broadcast.Broadcast, w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)

  if err != nil {
    log.Println(err)
    return
  }

  client := newClient(broadcast, conn)
  client.broadcast.RegisterClient(client)

  go client.BroadcastListener() 
  go client.ClientListener()
}
