package server

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/suphachok09790/hello-sekai-shop/pkg/response"
)

type healthCheck struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func (s *server) healthCheckService(c echo.Context) error {
	return response.SuccessResponse(c, http.StatusOK, &healthCheck{
		App:    s.cfg.App.Name,
		Status: "OK",
	})
}
