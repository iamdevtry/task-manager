package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/component"
	"github.com/iamdevtry/task-manager/db/model"
	"github.com/iamdevtry/task-manager/db/query"
)

func AddActivity(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var activity model.ActivityCreate

		err := ctx.ShouldBindJSON(&activity)

		if err != nil {
			panic(err)
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		activity.UserId = requester.GetUserId()

		store := query.NewStore(appCtx.GetDBConn())

		if err := store.CreateActivity(ctx.Request.Context(), activity); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}

func ListActivity(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()

		store := query.NewStore(appCtx.GetDBConn())

		activities, err := store.ListActivityByUser(ctx.Request.Context(), userId)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(activities))
	}
}
