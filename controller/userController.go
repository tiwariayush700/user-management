package controller

import "github.com/tiwariayush700/user-management/repository"

type userController struct {
	repository repository.UserRepository
	app        *app
}

func NewUserController(repository repository.UserRepository, app *app) *userController {
	return &userController{repository: repository, app: app}
}

func (u *userController) RegisterRoutes() {
	router := u.app.Router
	_ = router.Group("/")
	{
		//routerGroup.POST("/login", controller.Login())
		//
		//routerGroupVerified := routerGroup.Use(VerifyUserAndServe(controller.authService))
		//routerGroupVerified.GET("/me", controller.GetUserProfile())
	}
}
