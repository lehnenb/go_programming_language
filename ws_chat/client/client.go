package client

import (
  "bytes"
  "log"
  "fmt"
  "time"
  "github.com/gorilla/websocket"
  "github.com/lehnenb/go_programming_language/ws_chat/broadcast"
)

const (
  // Time allowed to write a message to the peer.
  writeWait = 10 * time.Second

  // Time allowed to read the next pong message from the peer.
  pongWait = 60 * time.Second

  // Send pings to peer with this period. Must be less than pongWait.
  pingPeriod = (pongWait * 9) / 10

  // Maximum message size allowed from peer.
  maxMessageSize = 512
)

var (
  newLine = []byte{'\n'}
  space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

// Client represents the connection between an individual user and the broadcast.
type Client struct {
  broadcast *broadcast.Broadcast

  // The websocket connection.
  conn *websocket.Conn
}

func newClient(broadcast *broadcast.Broadcast, conn *websocket.Conn) Client {
  return Client{
    broadcast: broadcast,
    conn: conn,
  }
}

// ClientListener sends messages from the websocket connection to the broadcast.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ClientListener() {
  defer func() {
    c.broadcast.Close()
    c.conn.Close()
  }()
  c.conn.SetReadLimit(maxMessageSize)
  c.conn.SetReadDeadline(time.Now().Add(pongWait))
  c.conn.SetPongHandler(func(string) error {
    c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil
  })

  for {
    _, message, err := c.conn.ReadMessage()

    if err != nil {
      if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
        log.Printf("error: %v", err)
      }
      break
    }

    message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
    c.broadcast.Publish(string(message))
  }
}

// BroadcastListener pumps messages from the broadcast to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) BroadcastListener() {
  ticker := time.NewTicker(pingPeriod)

  defer func() {
    ticker.Stop()
    c.conn.Close()
  }()

  for {
    select {
    case message, ok := <-c.broadcast.PubSub.Channel():
      c.conn.SetWriteDeadline(time.Now().Add(writeWait))
      if !ok {
        // The broadcast closed the channel.
        c.conn.WriteMessage(websocket.CloseMessage, []byte{})
        return
      }

      w, err := c.conn.NextWriter(websocket.TextMessage)
      if err != nil {
        return
      }
      w.Write([]byte(message.String()))
      fmt.Println(message.String())

      if err := w.Close(); err != nil {
        return
      }
    case <-ticker.C:
      c.conn.SetWriteDeadline(time.Now().Add(writeWait))
      if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
        return
      }
    }
  }
}
