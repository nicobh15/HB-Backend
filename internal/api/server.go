package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", server.createUser)
		userRoutes.GET("/:username", server.fetchUserByEmail)
		userRoutes.GET("/", server.listUsers)
		userRoutes.GET("/household", server.listUsersByHousehold)
		userRoutes.DELETE("/:email", server.deleteUser)
		userRoutes.PUT("/", server.updateUser)
	}

	householdRoutes := router.Group("/households")
	{
		householdRoutes.POST("/", server.createHousehold)
		householdRoutes.GET("/household", server.fetchHousehold)
		householdRoutes.GET("/", server.listHouseholds)
		householdRoutes.DELETE("/", server.deleteHousehold)
		householdRoutes.PUT("/", server.updateHousehold)
	}

	inventoryRoutes := router.Group("/inventory")
	{
		inventoryRoutes.POST("/", server.createInventoryItem)
		inventoryRoutes.GET("/", server.fetchHouseholdInventory)
		inventoryRoutes.DELETE("/", server.deleteInventoryItem)
		inventoryRoutes.PUT("/", server.updateInventoryItem)
	}

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
