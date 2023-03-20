package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/component"
	"github.com/iamdevtry/task-manager/db/model"
	"github.com/iamdevtry/task-manager/db/query"
)

func AddTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task model.TaskCreate

		err := ctx.ShouldBindJSON(&task)

		if err != nil {
			panic(err)
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		task.UserId = requester.GetUserId()

		store := query.NewStore(appCtx.GetDBConn())

		if err := store.CreateTask(ctx.Request.Context(), task); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}

func ListTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		store := query.NewStore(appCtx.GetDBConn())

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()

		tasks, err := store.GetTaskByUserId(ctx.Request.Context(), userId)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(tasks, nil, nil))
	}
}

// Not required - just a side effect
// func CountTask(appCtx component.AppContext) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		store := query.NewStore(appCtx.GetDBConn())

// 		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
// 		userId := requester.GetUserId()

// 		count, err := store.CountTaskByUserId(ctx.Request.Context(), userId)

// 		if err != nil {
// 			panic(err)
// 		}

// 		ctx.JSON(http.StatusOK, common.NewSuccessResponse(count, nil, nil))
// 	}
// }
