package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tiwariayush700/user-management/config"
	"github.com/tiwariayush700/user-management/models"
	repositoryImpl "github.com/tiwariayush700/user-management/repository/impl"
	serviceImpl "github.com/tiwariayush700/user-management/services/impl"
	"gorm.io/gorm"
)

// App structure for tenant microservice
type app struct {
	Config *config.Config
	DB     *gorm.DB //set from main.go
	Router *gin.Engine
}

func NewApp(config *config.Config, db *gorm.DB, router *gin.Engine) *app {
	return &app{
		Config: config,
		DB:     db,
		Router: router,
	}
}

func (app *app) Start() {

	app.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "HEAD", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	//repositories
	userRepositoryImpl := repositoryImpl.NewUserRepositoryImpl(app.DB)

	//services
	userService := serviceImpl.NewUserServiceImpl(userRepositoryImpl)

	//controllers
	userController := NewUserController(userService, app)

	//register routes
	userController.RegisterRoutes()

	app.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	logrus.Info("Application loaded successfully ")
	logrus.Fatal(app.Router.Run(":" + app.Config.Port))

}

func (app *app) Migrate() error {
	if err := app.DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	return nil
}
