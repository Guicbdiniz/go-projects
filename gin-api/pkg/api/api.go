package api

import (
	"github.com/Guicbdiniz/go-projects/gin-api/internal/routes/ping"
	"github.com/gin-gonic/gin"
)

func CreateAPI() (*gin.Engine, error) {
	api := gin.New()

	api.Use(gin.Logger())

	ping.AddPingRoute(api)

	return api, nil
}
