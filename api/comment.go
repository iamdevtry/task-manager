package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/component"
	"github.com/iamdevtry/task-manager/db/model"
	"github.com/iamdevtry/task-manager/db/query"
)

func AddComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var comment model.CommentCreate

		err := ctx.ShouldBindJSON(&comment)

		if err != nil {
			panic(err)
		}

		store := query.NewStore(appCtx.GetDBConn())

		result := &model.Comment{}
		if result, err = store.AddComment(ctx.Request.Context(), comment); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

func ListCommentsByActivityId(aptCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activityId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		store := query.NewStore(aptCtx.GetDBConn())

		comments, err := store.ListCommentsByActivityId(ctx.Request.Context(), int64(activityId))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(comments))
	}
}

func DeleteComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		store := query.NewStore(appCtx.GetDBConn())

		if err := store.DeleteComment(ctx.Request.Context(), int64(id)); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
