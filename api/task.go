package api

import (
	"fmt"
	"net/http"
	"strconv"

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

func DeleteTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		taskId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()

		store := query.NewStore(appCtx.GetDBConn())
		task, err := store.GetTask(ctx.Request.Context(), int64(taskId))
		if err != nil {
			panic(err)
		}
		if task.UserId != userId {
			panic(common.ErrCannotDeletedEntity("task", fmt.Errorf("You are not the owner of this task")))
		}

		if err := store.DeleteTask(ctx.Request.Context(), int64(taskId)); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}

func UpdateTask(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		taskId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}
		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()

		store := query.NewStore(appCtx.GetDBConn())
		task, err := store.GetTask(ctx.Request.Context(), int64(taskId))
		if err != nil {
			panic(err)
		}

		if task.UserId != userId {
			panic(common.ErrCannotUpdatedEntity("task", fmt.Errorf("You are not the owner of this task")))
		}
		var oldTask model.TaskUpdate
		err = ctx.ShouldBindJSON(&oldTask)
		if err != nil {
			panic(err)
		}

		if err := store.UpdateTask(ctx.Request.Context(), int64(taskId), oldTask); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
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
