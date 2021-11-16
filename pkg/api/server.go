package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var (
	logger = logrus.WithFields(logrus.Fields{"service": "api"})
)

func registerHandlers(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."+"/static/"))))
	router.HandleFunc("/", indexHandler).HeadersRegexp("User-Agent", "^Mozilla.*").Methods("GET")
	router.HandleFunc("/", infoHandler).Methods("GET")
	router.HandleFunc("/welcome", welcomeHandler).Methods("GET")
	router.HandleFunc("/widgets", widgetHandler)
	router.HandleFunc("/widgets/{id}", widgetHandler)
}

// ListenAndServe starts server and begins taking requests
func ListenAndServe() {
	logger.Infof("Starting server...")

	router := mux.NewRouter()

	// create routes
	registerHandlers(router)

	// create http server
	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}

	logger.Fatal(srv.ListenAndServe())
}
