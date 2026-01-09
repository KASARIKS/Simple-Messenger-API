package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kasariks/simple_messenger_api/services/message"
	"github.com/kasariks/simple_messenger_api/services/user"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	router *http.ServeMux
}

func NewServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr:   addr,
		db:     db,
		router: http.NewServeMux(),
	}
}

func (s *APIServer) Run() error {
	s.registerServiceRoutes()

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, s.router)
}

func (s *APIServer) registerServiceRoutes() {
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(s.router)

	messageStore := message.NewStore(s.db)
	messageHandler := message.NewHandler(messageStore, userStore)
	messageHandler.RegisterRoutes(s.router)
}
