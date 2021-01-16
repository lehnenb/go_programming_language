package main

import (
	"context"
	"sync"
	"github.com/go-redis/redis/v8"
)

// Broadcast contains the broadcast's connection data
type Broadcast struct {
  // mu Mutex for coordinating writes to the broadcasts map
  mu sync.Mutex

  // ID of the broadcast
  id string

  // Name of the broadcast
  name string

  // Redis PubSub connection
  pubSub *redis.PubSub

  // Closed indicates whether the broadcast has been closed
  closed bool
}

var broadcasts = make(map[string]*Broadcast)

func newBroadcast(name string, rClient *redis.Client) *Broadcast {
  ctx := context.Background()

  return &Broadcast{
    name: name,
    pubSub: rClient.Subscribe(ctx, name),
  }
}

// Close persists and then closes the broadcast
func (b *Broadcast) Close() {
  if (b.closed) {
    return;
  }

  b.mu.Lock()
  defer b.mu.Unlock()

  b.pubSub.Close()
  b.closed = true

  delete(broadcasts, b.name)
}
