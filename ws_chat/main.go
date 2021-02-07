package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/lehnenb/go_programming_language/ws_chat/broadcast"
	"github.com/lehnenb/go_programming_language/ws_chat/client"
)

var defaultBroadcast *broadcast.Broadcast
var redisClient *redis.Client
var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
  log.Println(r.URL)

  if r.URL.Path != "/" {
    http.Error(w, "Not found", http.StatusNotFound)
    return
  }

  if r.Method != "GET" {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  http.ServeFile(w, r, "home.html")
}


func init() {
  redisClient = redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })

  defaultBroadcast = broadcast.NewBroadcast("testarossa", redisClient)
}

func main() {
  flag.Parse()

  go defaultBroadcast.ListenToBroadcast()

  http.HandleFunc("/", serveHome)

  http.HandleFunc("/broadcast", func(w http.ResponseWriter, r *http.Request) {
    // createBroadcast(w, r)
  })

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    client.ServeClient(defaultBroadcast, w, r)
  })

  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
