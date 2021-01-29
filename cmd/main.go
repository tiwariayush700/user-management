package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	application "github.com/tiwariayush700/user-management/app"
	"github.com/tiwariayush700/user-management/config"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	appConfig := config.GetAppConfiguration()

	router := gin.Default()

	app := application.NewApp(appConfig, router)

	err := app.Migrate()
	if err != nil {
		logrus.Fatalf("err migrating db with err : %v", err)
	}

	app.Start()

}
