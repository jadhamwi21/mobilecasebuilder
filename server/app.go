package main

import (
	"github.com/jadhamwi21/mobilecasebuilder/config/middlewares"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	middlewares.ConfigPreMiddlewares(app)

	app.Start(":8080")

}
