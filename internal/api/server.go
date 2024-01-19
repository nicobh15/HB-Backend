package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.fetchUserByEmail)
	router.GET("/users", server.listUsers)
	router.GET("/users/household", server.listUsersByHousehold)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}