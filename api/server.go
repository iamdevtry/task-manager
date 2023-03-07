package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jmoiron/sqlx"
)

type Server struct {
	store  *db.DB
	router *gin.Engine
}

func NewServer(store *db.DB) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/accounts", server.listUsers)
	router.GET("/accounts/:id", server.getUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
