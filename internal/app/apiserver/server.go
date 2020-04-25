package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/seblex/rest-api-go/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

func newServer(store store.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
}

func (s *Server) handleUsersCreate() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}