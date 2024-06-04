package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/pusupalahemanthkumar/bankingsystem/db/sqlc"
	"github.com/pusupalahemanthkumar/bankingsystem/token"
	"github.com/pusupalahemanthkumar/bankingsystem/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()

	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Add routes to router
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
