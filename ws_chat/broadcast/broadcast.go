package broadcast

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lehnenb/go_programming_language/ws_chat/client"
)

const clientsPerBroadcast = 50

// Broadcast contains the broadcast's connection data
type Broadcast struct {
  // mu Mutex for coordinating writes to the broadcasts map
  mu sync.Mutex

  // Ctx carries the broadcasts main context
  Ctx context.Context

  // ID of the broadcast
  ID string

  // Name of the broadcast
  Name string

  // Redis pubSub connection
  pubSub *redis.PubSub

  // redisClient stores the reference for the redis client
  redisClient *redis.Client

  // clients stores references to subscribed clients
  clients []*client.Client

  // cancel
  cancel context.CancelFunc
}

type genericClient interface {
  Write(msg []byte)
}

var broadcasts = make(map[string]*Broadcast)

// NewBroadcast creates a new struct of the Broadcast type
func NewBroadcast(name string, rClient *redis.Client) *Broadcast {
  ctx, cancel := context.WithCancel(rClient.Context())

  return &Broadcast{
    Ctx: ctx,
    ID: fmt.Sprintf("%d", time.Now().UnixNano()),
    Name: name,
    redisClient: rClient,
    cancel: cancel,
    clients: make([]*client.Client, 0, clientsPerBroadcast),
  }
}

// ListenToBroadcast listens for message on the redis channel and publishes to all the subscribers
func (b *Broadcast) ListenToBroadcast() {
  ch := b.redisClient.Subscribe(context.Background(), b.Name).Channel()

  for {
    select {
    case msg, ok := <-ch:
      if(!ok) {
        fmt.Print("deu ruim")
        return;
      }

      fmt.Println(fmt.Sprintf("Write to clients: %s", msg))
      b.writeToClients([]byte(msg.Payload))
    }
  }
}

func (b *Broadcast) writeToClients(msg []byte) error {
  // Todo: implement retry mechanism
  for _, client := range b.clients {
    (*client).Write(msg)
  }

  return nil
}

// RegisterClient publishes and stores the register client records
func (b *Broadcast) RegisterClient(cli *client.Client) error {
  if (len(b.clients) == clientsPerBroadcast) {
    errMsg := fmt.Sprintf("Broadcast reached its full capacity")
    return errors.New(errMsg)
  }

  b.clients = append(b.clients, cli)

  return b.redisClient.Do(
    b.Ctx,
    "PUBLISH",
    b.Name,
    "User has entered the broadcast",
  ).Err()
}

// UnregisterClient publishes message about user leaving the broadcast
func (b *Broadcast) UnregisterClient(cli *client.Client) error {
  for index, client := range b.clients {
    if (client == cli) {
      blen := len(b.clients)
      b.clients[index] = b.clients[blen-1] // Copy last element to index i.
      b.clients[blen-1] = nil   // Erase last element (write zero value).
      b.clients = b.clients[:blen-1]   // Truncate slice.
    }
  }

  return b.redisClient.Do(
    b.Ctx,
    "PUBLISH",
    fmt.Sprintf("%s:%s", b.ID, b.Name),
    "User has left the broadcast",
  ).Err()
}

// Publish Publishes a message in the broadcast
func (b *Broadcast) Publish(message string) error {
  return b.redisClient.Do(
    b.Ctx,
    "PUBLISH",
    b.Name,
    message,
  ).Err()
}

// Close persists and then closes the broadcast
func (b *Broadcast) Close() {
  b.mu.Lock()
  defer b.mu.Unlock()

  b.pubSub.Close()
  delete(broadcasts, b.Name)

  if (b.Ctx.Err() == nil) {
    b.cancel()
  }
}
