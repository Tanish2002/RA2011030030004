package handlers

import (
	"net/http"
	"strconv"
	"train-service/controllers"

	"github.com/labstack/echo/v4"
)
 
type Handler struct {
	Controller controllers.Controllers
}

func (h *Handler) GetTrains(c echo.Context) error {
	trains := h.Controller.GetTrains()
	return c.JSON(http.StatusOK, trains)
}

func (h *Handler) GetTrain(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id required")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id should be a number")
	}
	trains := h.Controller.GetTrain(idInt)
	return c.JSON(http.StatusOK, trains)
}
