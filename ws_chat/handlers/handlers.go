package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-redis/redis/v8"
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
func NewApp(chr *chrome.Chrome) *App {
  app := App{chrome: chr}

  httpRouter := mux.NewRouter()
  httpRouter.HandleFunc("/export", app.export).Methods("POST")

  app.handler = httpRouter

  return &app
}

func (app *App) export(w http.ResponseWriter, r *http.Request) {
  reqLogger := logger.NewLogger()

  reqLogger.Log("request", "started")
  reqLogger.Log("request-method", r.Method)
  body, err := ioutil.ReadAll(r.Body)

  if len(body) == 0 {
    reqLogger.Log("request-error", "empty request")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if err != nil {
    reqLogger.Log("request-error", err.Error())
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  pdfData, err := app.chrome.ExportToPDF(&body)

  if err != nil {
    reqLogger.Log("request-error", err.Error())
  }

  _, err = w.Write(*pdfData)

  if err != nil {
    reqLogger.Log("request-error", err.Error())
    w.WriteHeader(http.StatusBadRequest)
  } else {
    reqLogger.Log("request-success", "PDF converted successfully")
    w.Header().Add("Content-Type", "application/pdf")
  }
}
