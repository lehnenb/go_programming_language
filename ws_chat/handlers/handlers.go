package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-redis/redis/v8"
	"github.com/lehnenb/go_programming_language/ws_chat/logger"
)

// App contains the data of a running application
// App.chrome contains the main browser context (and process)
// App.handler contains the application's HTTP router with all its routes and handlers
type App struct {
  redisClient *redis.Client
  handler http.Handler
}

// ServerHTTP implements the http.Handler interface into the App type.
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  app.handler.ServeHTTP(w, r)
}

// NewApp creates a new App instance
func NewApp(redisClient *redis.Client) *App {
  app := App{redisClient: redisClient}

  httpRouter := mux.NewRouter()
  httpRouter.HandleFunc("/client/ws", app.serveClientConnection).Methods("GET")

  app.handler = httpRouter

  return &app
}

func (app *App) serveClientConnection(w http.ResponseWriter, r *http.Request) {
  reqLogger := logger.NewLogger()

  reqLogger.Log("request", "started")
  reqLogger.Log("request-method", r.Method)

  if err != nil {
    reqLogger.Log("request-error", err.Error())
    w.WriteHeader(http.StatusBadRequest)
    return
  }


}
