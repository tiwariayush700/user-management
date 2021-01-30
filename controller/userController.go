package controller

import (
	"github.com/gin-gonic/gin"
	userError "github.com/tiwariayush700/user-management/error"
	"github.com/tiwariayush700/user-management/models"
	"github.com/tiwariayush700/user-management/services"
	"net/http"
)

type userController struct {
	service services.UserService
	app     *app
}

func NewUserController(service services.UserService, app *app) *userController {
	return &userController{service: service, app: app}
}

func (u *userController) RegisterRoutes() {
	router := u.app.Router
	userRouterGroup := router.Group("/user")
	{
		userRouterGroup.POST("/register", u.Register())
		userRouterGroup.POST("/login")

		//routerGroupVerified := routerGroup.Use(VerifyUserAndServe(controller.authService))
		//routerGroupVerified.GET("/me", controller.GetUserProfile())
	}
}

func (u *userController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		user := &models.User{}
		err := c.ShouldBind(user)
		if err != nil {
			err = userError.NewErrorBadRequest("invalid input")
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = u.service.CreateUser(c, user)
		if err != nil {
			err = userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		//todo return jwt token
		c.JSON(http.StatusOK, user)
	}
}
