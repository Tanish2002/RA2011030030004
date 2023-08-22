package main

import (
	"train-service/controllers"
	"train-service/handlers"
	"train-service/services"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	handler := handlers.Handler{
		Controller: controllers.Controllers{Service: *services.NewToken()},
	}

	e.GET("/schedules", handler.GetTrains)
	e.GET("/schedules/:id", handler.GetTrain)
	e.Logger.Fatal(e.Start(":8080"))
}
