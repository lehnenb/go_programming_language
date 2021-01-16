package main

import (
	"flag"
	"log"
	"net/http"
  "time"
	"github.com/go-redis/redis/v8"
)

var server  *http.Server
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
  server = &http.Server{
    Addr:         "0.0.0.0:3000",
    WriteTimeout: time.Second * 15,
    ReadTimeout:  time.Second * 15,
    IdleTimeout:  time.Second * 60,
    redisClient: redis.NewClient(&redis.Options{
      Addr: "localhost:6379",
      Password: "",
      DB: 0,
    }),
    Handler: app,
  }
}

func main() {
  flag.Parse()

  http.HandleFunc("/", serveHome)

  http.HandleFunc("/broadcast", func(w http.ResponseWriter, r *http.Request) {
    createBroadcast(w, r)
  })

  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    serveClient(hub, w, r)
  })

  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
