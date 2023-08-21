package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rayato159/neversitup-kafka-lab/src/config"
	"github.com/Rayato159/neversitup-kafka-lab/src/models"
	"github.com/Rayato159/neversitup-kafka-lab/src/pkg/queue"
	"github.com/Rayato159/neversitup-kafka-lab/src/pkg/request"
	"github.com/Rayato159/neversitup-kafka-lab/src/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	ControllerHandler interface {
		HealthCheck(c echo.Context) error
		OrderProcessor(c echo.Context) error
	}

	controller struct {
		cfg *config.Config
	}
)

func NewControllerHandler(cfg *config.Config) ControllerHandler {
	return &controller{
		cfg: cfg,
	}
}

func (con *controller) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, models.ErrResponse{
		Message: "OK",
	})
}

func (con *controller) OrderProcessor(c echo.Context) error {
	wrapper := request.ContextWrapper(c)

	req := new(models.Order)

	if err := wrapper.Bind(req); err != nil {
		log.Println("Error: bind:", err.Error())
		return response.ErrResponse(c, http.StatusInternalServerError, "Error: order is invalid format")
	}

	reqInBytes, err := json.Marshal(req)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, "Error: message into bytes failed")
	}
	if err := queue.PushMessageToQueue([]string{con.cfg.Kafka.Url}, "order", reqInBytes); err != nil {
		log.Println("Error:", err.Error())
		return response.ErrResponse(c, http.StatusInternalServerError, "Error: push message failed")
	}

	return response.ErrResponse(c, http.StatusOK, "Success: Message has been pushed")
}
