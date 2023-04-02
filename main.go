package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/task-manager/api"
	"github.com/iamdevtry/task-manager/component"
	"github.com/iamdevtry/task-manager/middleware"
	config "github.com/iamdevtry/task-manager/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

func runService(db *sqlx.DB, secretKey string) error {
	appCtx := component.NewAppContext(db, secretKey)

	route := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	route.Use(cors.New(config), middleware.Recover(appCtx))

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := route.Group("/v1")

	v1.POST("/register", api.Register(appCtx))
	v1.POST("/login", api.Login(appCtx))

	users := v1.Group("/users")
	{
		users.GET("", api.ListUsers(appCtx))
		users.GET("/:id", api.GetUser(appCtx))
	}

	tasks := v1.Group("/tasks", middleware.RequireAuth(appCtx))
	{
		tasks.POST("", api.AddTask(appCtx))
		tasks.GET("", api.ListTask(appCtx))
		tasks.DELETE("/:id", api.DeleteTask(appCtx))
		tasks.PUT("/:id", api.UpdateTask(appCtx))
	}

	activities := v1.Group("/activities", middleware.RequireAuth(appCtx))
	{
		activities.POST("", api.AddActivity(appCtx))
		activities.GET("", api.ListActivity(appCtx))
		activities.GET("/:id", api.GetActivity(appCtx))
		activities.DELETE("/:id", api.DeleteActivity(appCtx))
		activities.PUT("", api.UpdateStatusActivity(appCtx))
		activities.PUT("/:id", api.UpdateActivity(appCtx))
	}

	comments := v1.Group("/comments", middleware.RequireAuth(appCtx))
	{
		comments.POST("", api.AddComment(appCtx))
		comments.GET("/:id", api.ListCommentsByActivityId(appCtx))
		comments.DELETE("/:id", api.DeleteComment(appCtx))
	}

	return route.Run()
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:  ", err)
	}

	connectionString := fmt.Sprintf("oracle://%s:%s@%s/%s", config.DBUsername, config.DBPassword, config.DBServer, config.DBService)

	db, err := sql.Open(config.DBDriver, connectionString)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	slqx := sqlx.NewDb(db, config.DBDriver)

	if err := runService(slqx, config.SysSecretKey); err != nil {
		log.Fatal("cannot run service:", err)
	}
}
