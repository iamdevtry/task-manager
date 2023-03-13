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

		err = store.CreateTask(ctx.Request.Context(), task)

		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
