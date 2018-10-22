package backend

import (
	"log"
	"net/http"
	"os"

	"github.com/smjn/ipl18/backend/dao"
	"github.com/smjn/ipl18/backend/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	uDao      dao.UserDAO
)

var SetupAndGetRouter = func() http.Handler {
	log.Println("Setting up routes...")
	r := mux.NewRouter()
	setupRoutes(r)

	//wrap in route logger
	return setupLogging(r)
}

func setupRoutes(r *mux.Router) {
	//handle ping
	r.PathPrefix("/pub/ping").Methods("GET").Handler(handler.PingHandler)

	apiRouter := r.PathPrefix("/api").Subrouter()
	setupApi(apiRouter)
}

func setupApi(r *mux.Router) {
	r.Handle("/users/",handler.UserGetHandler{}).Methods("GET")
	r.Handle("/users/{inumber}", handler.UserGetHandler{}).Methods("GET")
	r.Handle("/users/{inumber}", handler.UserPutHandler{}).Methods("PUT")
	r.Handle("/users/",handler.UserPostHandler{}).Methods("POST")
	r.Handle("/users/{inumber}",handler.UserDeleteHandler{}).Methods("DELETE")
}

func setupLogging(r http.Handler) http.Handler {
	//provided by gorilla mux
	return handlers.LoggingHandler(os.Stdout, r)
}

func init() {
	uDao = dao.UserDAO{}
}
