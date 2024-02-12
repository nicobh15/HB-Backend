package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/nicobh15/HomeBuddy-Backend/internal/db/sqlc"
	"github.com/nicobh15/HomeBuddy-Backend/internal/token"
	"github.com/nicobh15/HomeBuddy-Backend/internal/util"
)

// TODO - Add recipes API
// TODO - Add tests for all the API methods
// TODO - Add validation for all the API methods
// TODO - Add authentication middleware
// TODO - Add logging middleware

type Server struct {
	config     util.Config
	store      *db.SQLStore
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.SQLStore) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouters()

	return server, nil
}

func (server *Server) setupRouters() {
	router := gin.Default()
	router.POST("/", server.createUser)
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/")
	authRoutes.Use(authMiddleware(server.tokenMaker))
	{
		userRoutes := authRoutes.Group("/users")
		{
			userRoutes.GET("/:username", server.fetchUserByUserName)
			userRoutes.GET("/", server.listUsers)
			userRoutes.GET("/household", server.listUsersByHousehold)
			userRoutes.DELETE("/:email", server.deleteUser)
			userRoutes.PUT("/", server.updateUser)
		}

		householdRoutes := authRoutes.Group("/households")
		{
			householdRoutes.POST("/", server.createHousehold)
			householdRoutes.GET("/household", server.fetchHousehold)
			householdRoutes.GET("/", server.listHouseholds)
			householdRoutes.DELETE("/", server.deleteHousehold)
			householdRoutes.PUT("/", server.updateHousehold)
		}

		inventoryRoutes := authRoutes.Group("/inventory")
		{
			inventoryRoutes.POST("/", server.createInventoryItem)
			inventoryRoutes.GET("/", server.fetchHouseholdInventory)
			inventoryRoutes.DELETE("/", server.deleteInventoryItem)
			inventoryRoutes.PUT("/", server.updateInventoryItem)
		}
	}

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
