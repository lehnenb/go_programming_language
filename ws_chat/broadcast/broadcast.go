package broadcast

import (
	"context"
	"sync"
	"time"
  "fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lehnenb/go_programming_language/ws_chat/client"
)

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

  // Redis PubSub connection
  PubSub *redis.PubSub

  // redisClient stores the reference for the redis client
  redisClient *redis.Client

  // cancel 
  cancel context.CancelFunc
}

var broadcasts = make(map[string]*Broadcast)

// NewBroadcast creates a new struct of the Broadcast type
func NewBroadcast(name string, rClient *redis.Client) *Broadcast {
  ctx, cancel := context.WithCancel(rClient.Context())
  pubSub := rClient.Subscribe(ctx, name)

  return &Broadcast{
    Ctx: ctx,
    ID: string(time.Now().UnixNano()),
    Name: name,
    PubSub: pubSub,
    redisClient: rClient,
    cancel: cancel,
  }
}

// RegisterClient publishes and stores the register client records
func (b *Broadcast) RegisterClient(cli client.Client) {
  b.redisClient.Do(
    b.Ctx,
    "PUBLISH",
    fmt.Sprintf("%s:%s", b.ID, b.Name),
    "New client has entered the broadcast",
  )
}

// Publish Publishes a message in the broadcast
func (b *Broadcast) Publish(message string) (error) {
  cmd := b.redisClient.Do(
    b.Ctx,
    "PUBLISH",
    fmt.Sprintf("%s:%s", b.ID, b.Name),
    message,
  )

  _, err := cmd.Result()

  return err
}

// Close persists and then closes the broadcast
func (b *Broadcast) Close() {
  b.mu.Lock()
  defer b.mu.Unlock()

  b.PubSub.Close()
  delete(broadcasts, b.Name)

  if (b.Ctx.Err == nil) {
    b.cancel()
  }
}
