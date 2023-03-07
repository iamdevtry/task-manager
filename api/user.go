package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/db/query"
)

func (server *Server) listUsers(ctx *gin.Context) {
	store := query.NewStore(server.store)

	users, err := store.ListUsers(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server *Server) getUser(ctx *gin.Context) {
	store := query.NewStore(server.store)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := store.GetUser(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
