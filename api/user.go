package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/component"
	"github.com/iamdevtry/task-manager/db/model"
	"github.com/iamdevtry/task-manager/db/query"
)

func ListUsers(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		store := query.NewStore(appCtx.GetDBConn())

		users, err := store.ListUsers(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}

func GetUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		store := query.NewStore(appCtx.GetDBConn())

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			panic(err)
		}

		user, err := store.GetUser(ctx, id)

		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func CreateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.CreateUser

		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		store := query.NewStore(appCtx.GetDBConn())

		err = store.CreateUser(ctx, user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
