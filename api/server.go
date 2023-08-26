package api

import (

	db "github.com/bontusss/gobank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server {store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccounts)
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) (interface{}) {
    return gin.H{"error": err.Error()}
}