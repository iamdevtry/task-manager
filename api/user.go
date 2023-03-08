package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/common"
	"github.com/iamdevtry/task-manager/component"
	"github.com/iamdevtry/task-manager/component/hasher"
	"github.com/iamdevtry/task-manager/component/tokenprovider"
	"github.com/iamdevtry/task-manager/component/tokenprovider/jwt"
	"github.com/iamdevtry/task-manager/db/model"
	"github.com/iamdevtry/task-manager/db/query"
)

func ListUsers(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		store := query.NewStore(appCtx.GetDBConn())

		users, err := store.ListUsers(ctx)

		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(users))
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
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.UserCreate

		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			panic(err)
		}

		store := query.NewStore(appCtx.GetDBConn())
		hash := hasher.NewMd5Hash()
		user.Password = hash.Hash(user.Password)
		err = store.Register(ctx, user)

		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("User created successfully"))
	}
}

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userLogin model.UserLogin
		if err := ctx.ShouldBindJSON(&userLogin); err != nil {
			panic(err)
		}

		store := query.NewStore(appCtx.GetDBConn())
		user, err := store.Login(ctx, userLogin)
		if err != nil {
			panic(common.NewCustomError(err, "Invalid username or password", "401"))
		}

		hash := hasher.NewMd5Hash()
		userLogin.Password = hash.Hash(userLogin.Password)
		if user.PasswordHash != userLogin.Password {
			panic(common.NewCustomError(err, "Invalid username or password", "401"))
		}

		jwtToken := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		payload := tokenprovider.TokenPayload{
			UserId: int(user.Id),
		}

		accessToken, err := jwtToken.Generate(payload, 60*60*24*7)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(accessToken))
	}
}
