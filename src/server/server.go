package server

import (
	"github.com/Rayato159/neversitup-kafka-lab/src/config"
	"github.com/Rayato159/neversitup-kafka-lab/src/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	ServerHandler interface {
		Start()
	}

	server struct {
		cfg *config.Config
	}

	stockServer   struct{ *server }
	paymentServer struct{ *server }

	httpServer struct {
		app *echo.Echo
		*server
	}
)

func NewServer(cfg *config.Config, appName string) ServerHandler {
	switch appName {
	case "order":
		return &httpServer{
			app: echo.New(),
			server: &server{
				cfg: cfg,
			},
		}
	case "stock":
		return &stockServer{
			&server{
				cfg: cfg,
			},
		}
	case "payment":
		return &paymentServer{
			&server{
				cfg: cfg,
			},
		}
	}
	panic("Error: server type not found.")
}

func (s *httpServer) Start() {
	s.app.Use(middleware.Logger())

	c := controller.NewControllerHandler(s.cfg)

	s.app.GET("/", c.HealthCheck)
	s.app.POST("/order", c.OrderProcessor)

	s.app.Logger.Fatal(s.app.Start(s.cfg.App.Url))
}

func (s *stockServer) Start() {
	c := controller.NewKafkaController(s.cfg)
	c.StockProcessor()
}

func (s *paymentServer) Start() {
	c := controller.NewKafkaController(s.cfg)
	c.PaymentProcessor()
}
