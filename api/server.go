package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *sql.DB
	router *gin.Engine
}

func NewServer(store *sql.DB) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
